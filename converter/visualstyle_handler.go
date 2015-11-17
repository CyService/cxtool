package converter

type VisualStyleHandler struct {
}

func (vsHandler VisualStyleHandler) HandleAspect(aspect []interface{}) map[string]interface{} {

	vpCount := len(aspect)

	// Result Map
	attrMap := make(map[string]interface{})

	for i := 0; i < vpCount; i++ {
		vp := aspect[i].(VisualProperty)
		vp.PropertiesOf

		key := attr["n"].(string)
		attrMap[key] = attr["v"]
	}

	return attrMap
}
