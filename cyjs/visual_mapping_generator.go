package cyjs

import (
	"strings"
	"strconv"
	"sort"
	"regexp"
	"log"
)

const (
	entrySeparator = ","
	kvSeparator = "="

	// Cytoscape tags
	bgImg = "background-image"
	bgFit = "background-fit"

	// Supported by Passthrough Mapping
	nLabel = "NODE_LABEL"
	eLabel = "EDGE_LABEL"
	nWidth = "NODE_WIDTH"
	nHeight = "NODE_HEIGHT"
	eWidth = "EDGE_WIDTH"
	nFSize = "NODE_LABEL_FONT_SIZE"
	eFSize = "EDGE_LABEL_FONT_SIZE"
	cgPrefix = "NODE_CUSTOMGRAPHICS_"
)


type VisualMappingGenerator struct {
	VpConverter VisualPropConverter
}

var replaceInvalid = regexp.MustCompile(`^[^a-zA-Z_]+|[^a-zA-Z_0-9]+`)

/*
	Passthrough Mapping converter.
 */
func (vmGenerator VisualMappingGenerator) CreatePassthroughMapping(
vpName string, definition string, entry *SelectorEntry) {

	parts := strings.Split(definition, entrySeparator)
	if len(parts) != 2 {
		return
	}

	tagAndValue := strings.Split(parts[0], kvSeparator)

	if len(tagAndValue) != 2 {
		return
	}

	switch {
	case vpName == nLabel ||  vpName == eLabel:
		entry.CSS["content"] = "data(" + tagAndValue[1] + ")"
	case vpName == nWidth || vpName == eWidth:
		entry.CSS["width"] = "data(" + tagAndValue[1] + ")"
	case vpName == nHeight:
		entry.CSS["height"] = "data(" + tagAndValue[1] + ")"
	case vpName == nFSize || vpName == eFSize:
		entry.CSS["font-size"] = "data(" + tagAndValue[1] + ")"
	case strings.HasPrefix(vpName, cgPrefix):
		// Custom Graphics Mapping
		entry.CSS[bgImg] = "data(" + tagAndValue[1] + ")"
		entry.CSS[bgFit] = "cover"  // TODO: Support for other types
	}
}


/*
	Create selectors for each key-value pair of discrete mapping.
 */
func (vmGenerator VisualMappingGenerator) CreateDiscreteMappings(
vpName string, vpCytoscape string, definition string, selectorType string) []SelectorEntry {

	var mappings []SelectorEntry

	parts := strings.Split(definition, entrySeparator)
	entryLen := len(parts)

	if entryLen < 2 {
		return mappings
	}

	// Extract column and its type
	colName := strings.Split(parts[0], kvSeparator)
	typeName := strings.Split(parts[1], kvSeparator)

	// validate:
	if entryLen % 2 != 0 {
		// Invalid definition string.
		return mappings
	}

	for i := 2; i < entryLen; i = i + 2 {
		k := strings.Split(parts[i], kvSeparator)
		v := strings.Split(parts[i + 1], kvSeparator)

		colVal := k[2]
		vpVal := v[2]

		// Build selector string
		// Example: node[degree = 5]
		var selectorStr string

		// Replace invalid char
		colN := replaceInvalid.ReplaceAllString(colName[1], "_")
		log.Println(colN)


		if isNumberType(typeName[1]) {
			// ' is not necessary for numbers.
			selectorStr = selectorType + "[" + colN + " = " + colVal + "]"
		} else {
			selectorStr = selectorType + "[" + colN + " = '" + colVal + "']"
		}

		css := make(map[string]interface{})
		log.Println("#######################")
		log.Println( vpCytoscape + " = " + vpVal)

		mappedVal := vmGenerator.VpConverter.GetCyjsPropertyValue(vpCytoscape, vpVal)
		log.Println("Converted = " + mappedVal.(string))

		css[vpName] = mappedVal
		
		// This only happens when arrow color is locked.
		if vpCytoscape == "EDGE_UNSELECTED_PAINT" {
			css["source-arrow-color"] = mappedVal
			css["target-arrow-color"] = mappedVal
		}

		newSelector := SelectorEntry{Selector:selectorStr, CSS:css}
		mappings = append(mappings, newSelector)
	}
	return mappings
}

func arrowColorSync() {

}


/*
	Continuous Mapping Converter
 */
func (vmGenerator VisualMappingGenerator) CreateContinuousMappings(
vpName string, vpCytoscape string, vpDataType string, definition string, selectorType string) []SelectorEntry {

	var selectors []SelectorEntry

	parts := strings.Split(definition, entrySeparator)
	entryLen := len(parts)

	if entryLen < 2 {
		return selectors
	}

	// Validate: each Continuous Mapping Point has 4 entries.
	if (entryLen - 2) % 4 != 0 {
		return selectors
	}

	// Extract column and its type
	colName := strings.Split(parts[0], kvSeparator)
	typeName := strings.Split(parts[1], kvSeparator)

	columnName := colName[1]
	// Replace invalid char
	columnName = replaceInvalid.ReplaceAllString(colName[1], "_")

	columnDataType := typeName[1]

	// Assume all values are double in continuous mapping

	points := make(map[float64]interface{})

	for i := 2; i < entryLen; i = i + 4 {
		l := strings.Split(parts[i], kvSeparator)
		e := strings.Split(parts[i + 1], kvSeparator)
		g := strings.Split(parts[i + 2], kvSeparator)
		v := strings.Split(parts[i + 3], kvSeparator)

		ov, err := parseNumber(columnDataType, v[2])
		if err != nil {
			continue
		}

		point := make(map[string]interface{})
		point["l"] = l[2]
		point["e"] = e[2]
		point["g"] = g[2]
		point["ovString"] = v[2]

		points[ov.(float64)] = point

	}

	// Sort by key
	var keys []float64
	for k := range points {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	numPoints := len(points)
	if numPoints <= 0 {
		return selectors
	}


	if numPoints == 1 {
		// Case 1: only one point
	} else {
		// Special case: Size is not supported in Cytoscape.js.
		// Create two mappings instead.
		if vpCytoscape == "NODE_SIZE" {
			selectorsW := vmGenerator.multiplePointsMapping(points, keys,
				selectorType, columnName, vpDataType, "width", "NODE_WIDTH")
			selectorsH := vmGenerator.multiplePointsMapping(points, keys,
				selectorType, columnName, vpDataType, "height", "NODE_HEIGHT")
			selectors = append(selectors, selectorsW...)
			selectors = append(selectors, selectorsH...)
		} else {
			selectors = vmGenerator.multiplePointsMapping(points, keys,
				selectorType, columnName, vpDataType, vpName, vpCytoscape)
		}
	}

	return selectors
}

func (vmGenerator VisualMappingGenerator) multiplePointsMapping(points map[float64]interface{},
sortedKeys []float64, selectorType string,
columnName string, vpDataType string, vp string, vpCytoscape string) []SelectorEntry {

	numPoints := len(points)
	var selectors []SelectorEntry

	for idx, key := range sortedKeys {
		if idx == 0 {
			// First point
			p := points[key].(map[string]interface{})
			pStr := p["ovString"].(string)
			selectorLeft := selectorType + "[" + columnName + " < " + pStr + "]"
			selectorLeftEq := selectorType + "[" + columnName + " = " + pStr + "]"

			cssLeft := make(map[string]interface{})
			cssLeftEq := make(map[string]interface{})

			cssLeft[vp] = vmGenerator.VpConverter.GetCyjsPropertyValue(vpCytoscape, p["l"].(string))
			cssLeftEq[vp] = vmGenerator.VpConverter.GetCyjsPropertyValue(vpCytoscape, p["e"].(string))

			selectors = append(selectors, SelectorEntry{Selector:selectorLeft, CSS:cssLeft})
			selectors = append(selectors, SelectorEntry{Selector:selectorLeftEq, CSS:cssLeftEq})
		} else {

			var p, pNext map[string]interface{}

			if idx != (numPoints - 1) {
				p = points[key].(map[string]interface{})
				pNext = points[sortedKeys[idx + 1]].(map[string]interface{})
			} else {
				// This is the last point
				p = points[sortedKeys[idx - 1]].(map[string]interface{})
				pNext = points[sortedKeys[idx]].(map[string]interface{})
			}
			pStr := p["ovString"].(string)
			pStrNext := pNext["ovString"].(string)

			selectorMiddle := selectorType +
			"[" + columnName + " > " + pStr + "]" +
			"[" + columnName + " < " + pStrNext + "]"

			cssMiddle := make(map[string]interface{})
			s := []string{"mapData(", columnName, ",", pStr, ",", pStrNext, ",",
				p["g"].(string), ",", pNext["l"].(string), ")"}
			cssMiddle[vp] = strings.Join(s, "")
			selectors = append(selectors, SelectorEntry{Selector:selectorMiddle,
				CSS:cssMiddle})

			selectorNextEq := selectorType + "[" + columnName + " = " + pStrNext + "]"
			cssNextEq := make(map[string]interface{})
			cssNextEq[vp] = vmGenerator.VpConverter.GetCyjsPropertyValue(vpCytoscape, pNext["e"].(string))

			selectors = append(selectors, SelectorEntry{Selector:selectorNextEq,
				CSS:cssNextEq})

			if idx == (numPoints - 1) {
				// Last point: Add extra selector
				selectorRight := selectorType + "[" + columnName + " > " +
				pStrNext + "]"
				cssRight := make(map[string]interface{})
				cssRight[vp] = vmGenerator.VpConverter.GetCyjsPropertyValue(vpCytoscape, pNext["g"].(string))
				selectors = append(selectors, SelectorEntry{Selector:selectorRight,
					CSS:cssRight})
			}
		}
	}
	return selectors
}


func isNumberType(colType string) bool {
	switch colType{
	case "double", "integer", "float":
		return true
	default:
		return false
	}
}

func parseNumber(colType string, value string) (num interface{}, err error) {
	switch colType {
	case "double", "float", "integer", "long":
		dVal, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return dVal, nil
		}
		return dVal, nil
	default:
	}
	return value, nil
}