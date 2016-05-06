package converter

import (
	cx "github.com/cyService/cxtool/cx"
	cyjs "github.com/cyService/cxtool/cyjs"
	"log"
	"strconv"
)

const (
	network string = "network"
	cxNodes string = "nodes"
	cxEdges string = "edges"
	nodesDefault string = "nodes:default"
	edgesDefault string = "edges:default"

	// VP tags
	dependencies string = "dependencies"
	properties string = "properties"
	propOf string = "properties_of"

	// Visual Property Dependency List
	nodeSizeLocked string = "nodeSizeLocked"
	nodeCustomGraphicsSizeSync string = "nodeCustomGraphicsSizeSync"
	arrowColorMatchesEdge string = "arrowColorMatchesEdge"
)

var lockRelated = make(map[string]bool)

func init() {
	lockRelated["NODE_SIZE"] = true
	lockRelated["NODE_HEIGHT"] = true
	lockRelated["NODE_WIDTH"] = true
	lockRelated["EDGE_TARGET_ARROW_UNSELECTED_PAINT"] = true
	lockRelated["EDGE_SOURCE_ARROW_UNSELECTED_PAINT"] = true
	lockRelated["EDGE_STROKE_UNSELECTED_PAINT"] = true
	lockRelated["EDGE_UNSELECTED_PAINT"] = true
}

type VisualStyleHandler struct {
	conversionTable        map[string]string

	typeTable              map[string]string

	visualMappingGenerator cyjs.VisualMappingGenerator
}

func (vsHandler VisualStyleHandler) HandleAspect(aspect []interface{}) map[string]interface{} {

	// Type converter
	//	vpConverter := VisualPropConverter{typeTable: vsHandler.typeTable}
	var converter cyjs.VisualPropConverter = vsHandler.visualMappingGenerator.VpConverter

	// This is the number of elements in a VP section.
	vpCount := len(aspect)

	// Result Map
	vpMap := make(map[string]interface{})

	// Temp Visual Style object,
	//  A map from selector name to actual props.
	style := make(map[string]cyjs.SelectorEntry)

	var selectors []cyjs.SelectorEntry
	var defaultsSelectors []cyjs.SelectorEntry

	for i := 0; i < vpCount; i++ {
		// Extract a new selector
		vp := aspect[i].(map[string]interface{})

		targetProperty := vp[propOf].(string)

		// Dependencies
		var depList map[string]interface{}
		deps, exist := vp[dependencies]
		if exist {
			depList = deps.(map[string]interface{})
		}

		log.Println(depList)

		// Check valid graph object (node/edge/net) ot not
		selectorTag, isDefaults := isValidProperty(targetProperty)

		if selectorTag == "" {
			continue
		}

		// This is the actual entry to be added
		entry := cyjs.SelectorEntry{}

		entry.Selector = selectorTag

		// Properties list: Usually default values
		cxProps := vp[properties].(map[string]interface{})

		css := make(map[string]interface{})

		// Handle special cases: VP Locks
		handleDependencies(depList, css, cxProps, converter)

		for key, value := range cxProps {
			_, isLockKey := lockRelated[key]

			if isLockKey {
				continue
			}

			ag, exists := vsHandler.conversionTable[key]

			if !exists {
				continue
			}

			// Handle special Visual Props
			if key == "NODE_LABEL_POSITION" {
				translateLabelPosition(css, cxProps, converter)
				continue
			}

			convertedValue := vsHandler.visualMappingGenerator.VpConverter.GetCyjsPropertyValue(key, value.(string))
			css[ag] = convertedValue
		}
		entry.CSS = css

		mappings, exists := vp[cx.Mappings]
		if exists {
			// Parse mapping entries
			visualMappings := vsHandler.createMappings(
				selectorTag, mappings.(map[string]interface{}), &entry)

			selectors = append(selectors, visualMappings...)
		}

		// Save for later use
		// This is necessary for
		style[selectorTag] = entry

		if isDefaults {
			defaultsSelectors = append(defaultsSelectors, entry)
		} else {
			selectors = append(selectors, entry)
		}

	}

	// Add selectors under "style" tab
	mergedSelector := append(defaultsSelectors, selectors...)
	vpMap["style"] = mergedSelector

	return vpMap
}

func handleDependencies(depList map[string]interface{},
css map[string]interface{}, cxProps map[string]interface{}, converter cyjs.VisualPropConverter) {

	// Check VP dependency list actually exists or not.
	if depList == nil {
		return
	}

	// Node size dependency
	sizeDep, exist := depList[nodeSizeLocked]
	if exist {
		processNodeSizeLocked(sizeDep.(string), css, cxProps, converter)
	}

	arrowDep, exist := depList[arrowColorMatchesEdge]
	if exist {
		processEdgeArrowColor(arrowDep.(string), css, cxProps, converter)
	}
}

func processNodeSizeLocked(sizeLockedStr string,
css map[string]interface{}, cxProps map[string]interface{}, converter cyjs.VisualPropConverter) {
	sizeLocked, _ := strconv.ParseBool(sizeLockedStr)

	log.Println("Size LOCK:")
	log.Println(sizeLocked)

	if sizeLocked {
		log.Println("THIS IS LOCKED:")
		value := cxProps["NODE_SIZE"]
		convertedValue := converter.GetCyjsPropertyValue("NODE_SIZE", value.(string))
		css["height"] = convertedValue
		css["width"] = convertedValue
	} else {
		w := cxProps["NODE_WIDTH"]
		h := cxProps["NODE_HEIGHT"]
		wValue := converter.GetCyjsPropertyValue("NODE_WIDTH", w.(string))
		hValue := converter.GetCyjsPropertyValue("NODE_HEIGHT", h.(string))
		css["height"] = hValue
		css["width"] = wValue
	}
}

func processEdgeArrowColor(arrowLockedStr string,
css map[string]interface{}, cxProps map[string]interface{}, converter cyjs.VisualPropConverter) {
	arrowLocked, _ := strconv.ParseBool(arrowLockedStr)

	log.Println("ARROW LOCK:")
	log.Println(arrowLocked)

	if arrowLocked {
		// Need to use EDGE_UNSELECTED_PAINT if locked.
		value := cxProps["EDGE_UNSELECTED_PAINT"]
		convertedValue := converter.GetCyjsPropertyValue("EDGE_UNSELECTED_PAINT", value.(string))
		css["target-arrow-color"] = convertedValue
		css["source-arrow-color"] = convertedValue
		css["line-color"] = convertedValue
	} else {
		l := cxProps["EDGE_STROKE_UNSELECTED_PAINT"]
		s := cxProps["EDGE_SOURCE_ARROW_UNSELECTED_PAINT"]
		t := cxProps["EDGE_TARGET_ARROW_UNSELECTED_PAINT"]
		lColor := converter.GetCyjsPropertyValue("EDGE_STROKE_UNSELECTED_PAINT", l.(string))
		sColor := converter.GetCyjsPropertyValue("EDGE_STROKE_UNSELECTED_PAINT", s.(string))
		tColor := converter.GetCyjsPropertyValue("EDGE_STROKE_UNSELECTED_PAINT", t.(string))
		css["line-color"] = lColor
		css["source-arrow-color"] = sColor
		css["target-arrow-color"] = tColor
	}

}

func translateLabelPosition(css map[string]interface{},
cxProps map[string]interface{}, converter cyjs.VisualPropConverter) {

	log.Println("LABEL POS: ")

	value := cxProps["NODE_LABEL_POSITION"]
	pos := converter.GetCyjsPropertyValue("NODE_LABEL_POSITION", value.(string)).([2]string)
	css["text-valign"] = pos[0]
	css["text-halign"] = pos[1]

}

func (vsHandler VisualStyleHandler) createMappings(selectorTag string,
mappings map[string]interface{}, entry *cyjs.SelectorEntry) (newSelectors []cyjs.SelectorEntry) {

	var newMaps []cyjs.SelectorEntry

	for vp, mapping := range mappings {
		visualMapping := mapping.(map[string]interface{})
		mappingType := visualMapping["type"].(string)
		definition := visualMapping["definition"].(string)

		switch mappingType {
		case cx.Passthrough:
			vsHandler.visualMappingGenerator.CreatePassthroughMapping(vp,
				definition, entry)
		case cx.Discrete:
			ag := vsHandler.conversionTable[vp]
			newMappings := vsHandler.visualMappingGenerator.CreateDiscreteMappings(
				ag, vp,
				definition, selectorTag)
			newMaps = append(newMaps, newMappings...)
		case cx.Continuous:
			ag := vsHandler.conversionTable[vp]
			vpDataType := vsHandler.typeTable[vp]
			newMappings := vsHandler.visualMappingGenerator.CreateContinuousMappings(ag, vp, vpDataType, definition, selectorTag)
			newMaps = append(newMaps, newMappings...)
		default:
		}
	}

	return newMaps
}


//
// Check the given "property_of" tag is valid or not.
// 2nd parameter is true if it is a list of defaults
//
func isValidProperty(propertyOf string) (tag string, defaults bool) {
	switch propertyOf {
	case nodesDefault:
		return cx.NodeTag, true
	case cxNodes:
		return cx.NodeTag, false
	case edgesDefault:
		return cx.EdgeTag, true
	case cxEdges:
		return cx.EdgeTag, false
	case network:
		return "", false
	default:
		return "", false
	}
}