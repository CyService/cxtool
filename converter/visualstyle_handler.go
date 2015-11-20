package converter

const (
	network      string = "network"
	cxNodes      string = "nodes"
	cxEdges      string = "edges"
	nodesDefault string = "nodes:default"
	edgesDefault string = "edges:default"
)

type VisualStyleHandler struct {
	conversionTable map[string]string

	typeTable map[string]string
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

		// Save for later use
		// This is necessary for
		style[selectorTag] = entry
		selectors = append(selectors, entry)
	}
	vpMap["style"] = selectors

	return vpMap
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
