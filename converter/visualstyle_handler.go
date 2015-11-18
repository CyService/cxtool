package converter

type VisualStyleHandler struct {
}

func (vsHandler VisualStyleHandler) HandleAspect(aspect []interface{}) map[string]interface{} {

	vpCount := len(aspect)

	// Result Map
	attrMap := make(map[string]interface{})

	for i := 0; i < vpCount; i++ {

		// Create new selector
		entry := SelectorEntry{}
		vp := aspect[i].(VisualProperty)
		entry.Selector = vp.PropertiesOf
	}

	return attrMap
}
