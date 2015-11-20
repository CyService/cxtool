package converter

import (
	"strconv"
	"reflect"
)

type AttributeHandler struct {
	typeDecoder TypeDecoder
}

func (attrHandler AttributeHandler) HandleAspect(aspect []interface{}) map[string]interface{} {

	// Find length of this aspects to be processed
	attrCount := len(aspect)
	values := make(map[string]interface{})

	for i := 0; i < attrCount; i++ {
		attr := aspect[i].(map[string]interface{})


		// Pointer value
		ptr := attr["po"]

		// Extract pointer and check type
		pointer := strconv.FormatInt(int64(ptr.(float64)), 10)

		// Check the value already exists or not
		attrMap, exist := values[pointer]

		attrEntries := make(map[string]interface{})
		if exist {
			attrEntries = attrMap.(map[string]interface{})
		}

		attributeName := attr["n"].(string)

		value := attr["v"]

		// This is optional (data type)
		dataType, exists := attr["d"]
		if exists && reflect.TypeOf(value) == reflect.TypeOf("") {
			// Need data type conversion
			value = attrHandler.typeDecoder.decode(value.(string), dataType.
			(string))
		}

		attrEntries[attributeName] = value

		values[pointer] = attrEntries
	}

	return values
}
