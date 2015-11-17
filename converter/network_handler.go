package converter

type NetworkHandler struct {
}

func (netHandler NetworkHandler) HandleAspect(aspect []interface{}) map[string]interface{} {

	attrCount := len(aspect)

	// Result Map
	attrMap := make(map[string]interface{})

	for i := 0; i < attrCount; i++ {
		attr := aspect[i].(map[string]interface{})
		key := attr["n"].(string)
		attrMap[key] = attr["v"]
	}

	return attrMap
}
