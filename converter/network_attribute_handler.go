package converter
import (
	cx "github.com/cytoscape-ci/cxtool/cx"
	"reflect"
)

type NetworkAttributeHandler struct {
	typeDecoder cx.TypeDecoder
}

func (networkAttrHandler NetworkAttributeHandler) HandleAspect(aspect []interface{}) map[string]interface{} {

	// Find length of this aspects to be processed
	attrCount := len(aspect)

	// Result Map
	attrMap := make(map[string]interface{})

	for i := 0; i < attrCount; i++ {
		attr := aspect[i].(map[string]interface{})
		processEntry(networkAttrHandler.typeDecoder, attr, attrMap)
	}

	return attrMap
}

func processEntry(decoder cx.TypeDecoder, attr map[string]interface{},
	attrMap map[string]interface{}) {
	key := attr["n"].(string)

	value := attr["v"].(interface{})

	dataType, exists := attr["d"]

	if exists && reflect.TypeOf(value) == reflect.TypeOf("") {
		// Need data type conversion
		value = decoder.Decode(value.(string), dataType.(string))
	}

	attrMap[key] = value
}
