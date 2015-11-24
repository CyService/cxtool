package converter
import (
	"strings"
	"strconv"
	"sort"
)


const (
	entrySeparator = ","
	kvSeparator = "="
)

type VisualMappingGenerator struct {
}


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

	// This mapping is valid only for Labels (at least foe now...)
	if vpName == "NODE_LABEL" || vpName == "EDGE_LABEL" {
		entry.CSS["contents"] = "data(" + tagAndValue[1] + ")"
	}
}


//
// Create selectors for each key-value pair of discrete mapping.
//
func (vmGenerator VisualMappingGenerator) CreateDiscreteMappings(
vpName string, definition string, selectorType string) []SelectorEntry {

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
	if entryLen%2 != 0 {
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

		if isNumberType(typeName[1]) {
			// ' is not necessary for numbers.
			selectorStr = selectorType + "[" + colName[1] + " = " + colVal + "]"
		} else {
			selectorStr = selectorType + "[" + colName[1] + " = '" + colVal + "']"
		}

		css := make(map[string]interface{})
		css[vpName] = vpVal

		newSelector := SelectorEntry{Selector:selectorStr, CSS:css}
		mappings = append(mappings, newSelector)
	}
	return mappings
}


func (vmGenerator VisualMappingGenerator) CreateContinuousMappings(
vpName string, definition string, selectorType string) []SelectorEntry {

	var selectors []SelectorEntry

	parts := strings.Split(definition, entrySeparator)
	entryLen := len(parts)

	if entryLen < 2 {
		return selectors
	}

	// Validate: each Continuous Mapping Point has 4 entries.
	if (entryLen-2)%4 != 0 {
		return selectors
	}

	// Extract column and its type
	colName := strings.Split(parts[0], kvSeparator)
	typeName := strings.Split(parts[1], kvSeparator)

	columnDataType := typeName[1]

	// Assume all values are double in continuous mapping


	points := make(map[float64]interface{})

	for i := 2; i < entryLen; i = i + 4 {
		l := strings.Split(parts[i], kvSeparator)
		e := strings.Split(parts[i+1], kvSeparator)
		g := strings.Split(parts[i+2], kvSeparator)
		v := strings.Split(parts[i+3], kvSeparator)

		ov, err := parseNumber(columnDataType, v[2])
		if err != nil {
			continue
		}

		lt, err := parseNumber(columnDataType, l[2])
		if err != nil {
			continue
		}
		eq, err := parseNumber(columnDataType, e[2])
		if err != nil {
			continue
		}
		gt, err := parseNumber(columnDataType, g[2])
		if err != nil {
			continue
		}

		point := make(map[string]interface{})
		point["l"] = lt
		point["e"] = eq
		point["g"] = gt
		points[ov.(float64)] = point

	}

	// Sort by key
	var keys []float64
	for k := range points {
    	keys = append(keys, k)
	}
	sort.Float64s(keys)

	for _, k := range keys {
		pt := points[k].(map[string]interface{})
		ptStr := strconv.FormatFloat(k, 'E', -1, 64)
		selectorStr := selectorType + "[" + colName[1] + " = " + ptStr + "]"

		newSelector := SelectorEntry{Selector:selectorStr, CSS:pt}
		selectors = append(selectors, newSelector)
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
	case "double", "float":
		dVal, err :=strconv.ParseFloat(value, 64)
		if err == nil {
			return dVal, nil
		}
	case "integer", "long":
		iVal, err :=strconv.ParseInt(value, 10, 64)
		if err == nil {
			return iVal, nil
		}
	default:
		return value, nil
	}
	return value, nil
}
