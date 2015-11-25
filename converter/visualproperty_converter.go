package converter

import (
	"strconv"
	"strings"
)

type VisualPropConverter struct {
	typeTable map[string]string
}

func (vpConverter VisualPropConverter) getCyjsPropertyValue(key string, value string) (converted interface{}) {

	dataType := vpConverter.typeTable[key]

	switch dataType {
	case "number":
		number, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return number
		} else {
			return value
		}
	case "font":
		return parseFont(value)
	default:
		return value
	}
}

func parseFont(value string) string {
	entries := strings.Split(value, ",")
	if len(entries) >= 1 {
		return entries[0]
	}

	return value
}
