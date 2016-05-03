package converter

import (
	cx "github.com/cytoscape-ci/cxtool/cx"
	cyjs "github.com/cytoscape-ci/cxtool/cyjs"
	"log"
	"strconv"
)

const (
	network      string = "network"
	cxNodes      string = "nodes"
	cxEdges      string = "edges"
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

type VisualStyleHandler struct {
	conversionTable        map[string]string

	typeTable              map[string]string

	visualMappingGenerator cyjs.VisualMappingGenerator
}

func (vsHandler VisualStyleHandler) HandleAspect(aspect []interface{}) map[string]interface{} {

	// Type converter
//	vpConverter := VisualPropConverter{typeTable: vsHandler.typeTable}

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

		// Handle Size:
		sizeLocked := false

		if depList != nil {
			_, exist := depList[nodeSizeLocked]
			if exist {

				sizeLockedStr := depList[nodeSizeLocked].(string)
				sizeLocked, _ = strconv.ParseBool(sizeLockedStr)

				log.Println("LOCK:")
				log.Println(sizeLocked)

				if sizeLocked {
					log.Println("THIS IS LOCKED:")
					value := cxProps["NODE_SIZE"]
					convertedValue := vsHandler.visualMappingGenerator.VpConverter.GetCyjsPropertyValue("NODE_SIZE", value.(string))
					css["height"] = convertedValue
					css["width"] = convertedValue
				} else {
					w := cxProps["NODE_WIDTH"]
					h := cxProps["NODE_HEIGHT"]
					wValue := vsHandler.visualMappingGenerator.VpConverter.GetCyjsPropertyValue("width", w.(string))
					hValue := vsHandler.visualMappingGenerator.VpConverter.GetCyjsPropertyValue("height", h.(string))
					css["height"] = hValue
					css["width"] = wValue
				}
			}
		}

		for key, value := range cxProps {
			if key == "NODE_WIDTH" || key == "NODE_HEIGHT" || key == "NODE_SIZE" {
				continue
			}

			ag, exists := vsHandler.conversionTable[key]

			if !exists {
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



func (vsHandler VisualStyleHandler) createMappings(selectorTag string,
mappings map[string]interface{}, entry *cyjs.SelectorEntry)(newSelectors []cyjs.SelectorEntry){

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
			newMappings := vsHandler.visualMappingGenerator.CreateDiscreteMappings(ag,
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