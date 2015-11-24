package converter

const (
	network      string = "network"
	cxNodes      string = "nodes"
	cxEdges      string = "edges"
	nodesDefault string = "nodes:default"
	edgesDefault string = "edges:default"
)

type VisualStyleHandler struct {
	conversionTable        map[string]string

	typeTable              map[string]string

	visualMappingGenerator VisualMappingGenerator
}

func (vsHandler VisualStyleHandler) HandleAspect(aspect []interface{}) map[string]interface{} {

	// Type converter
	vpConverter := VisualPropConverter{typeTable: vsHandler.typeTable}

	vpCount := len(aspect)

	// Result Map
	vpMap := make(map[string]interface{})

	// Temp Visual Style object,
	//  A map from selector name to actual props.
	style := make(map[string]SelectorEntry)

	var selectors []SelectorEntry

	for i := 0; i < vpCount; i++ {

		// Extract a new selector
		vp := aspect[i].(map[string]interface{})
		targetProperty := vp["properties_of"].(string)

		// Supported or not
		selectorTag := isValidProperty(targetProperty)

		if selectorTag == "" {
			continue
		}

		entry := SelectorEntry{}
		entry.Selector = selectorTag

		cxProps := vp["properties"].(map[string]interface{})

		css := make(map[string]interface{})

		for key, value := range cxProps {
			cyjsTag, exists := vsHandler.conversionTable[key]

			if !exists {
				continue
			}
			convertedValue := vpConverter.getCyjsPropertyValue(key, value.(string))
			css[cyjsTag] = convertedValue

		}
		entry.CSS = css

		mappings, exists := vp[cx_mappings]
		if exists {
			// Parse mapping entries
			visualMappings := vsHandler.createMappings(
				selectorTag, mappings.(map[string]interface{}), &entry)

			selectors = append(selectors, visualMappings...)
		}

		// Save for later use
		// This is necessary for
		style[selectorTag] = entry
		selectors = append(selectors, entry)
	}
	vpMap["style"] = selectors

	return vpMap
}

func (vsHandler VisualStyleHandler) createMappings(selectorTag string,
mappings map[string]interface{}, entry *SelectorEntry)(newSelectors []SelectorEntry){

	var newMaps []SelectorEntry

	for vp, mapping := range mappings {
		visualMapping := mapping.(map[string]interface{})
		mappingType := visualMapping["type"].(string)
		definition := visualMapping["definition"].(string)

		switch mappingType {
		case passthrough:
			vsHandler.visualMappingGenerator.CreatePassthroughMapping(vp,
				definition, entry)
		case discrete:
			cyjsTag := vsHandler.conversionTable[vp]
			newMappings := vsHandler.visualMappingGenerator.CreateDiscreteMappings(cyjsTag,
				definition, selectorTag)
			newMaps = append(newMaps, newMappings...)
		case continuous:
			cyjsTag := vsHandler.conversionTable[vp]
			newMappings := vsHandler.visualMappingGenerator.CreateContinuousMappings(cyjsTag, definition, selectorTag)
			newMaps = append(newMaps, newMappings...)
		default:
		}
	}

	return newMaps
}

func isValidProperty(propertyOf string) (tag string) {
	switch propertyOf {
	case nodesDefault:
		return node
	case cxNodes:
		return node
	case edgesDefault:
		return edge
	case cxEdges:
		return edge
	case network:
		return ""
	default:
		return ""
	}
}
