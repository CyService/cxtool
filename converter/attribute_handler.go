package converter

import (
	"strconv"
)
type AttributeHandler struct {
}

func (attrHandler AttributeHandler) HandleAspect(aspect[]interface{})map[string]interface{} {

	// Find length of this aspects to be processed
	attrCount := len(aspect)

	values := make(map[string]interface{})

	for i := 0; i < attrCount; i++ {
		attr :=aspect[i].(map[string]interface{})

		// Extract pointer (key)
		pointer := strconv.FormatInt(int64(attr["po"].(float64)), 10)

		// Check the value already exists or not
		attrMap, exist := values[pointer]

		attrEntries := make(map[string]interface{})
		if exist {
			attrEntries = attrMap.(map[string]interface{})
		}

		attributeName := attr["n"].(string)
		attrEntries[attributeName] = attr["v"]

		values[pointer] = attrEntries
	}

	return values
}
