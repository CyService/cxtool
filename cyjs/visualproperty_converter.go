package cyjs

import (
	"strconv"
	"strings"
)

type VisualPropConverter struct {
	typeTable map[string]string

	valueConverters map[string]ValueConverter
}

func NewVisualPropConverter(typeTable map[string]string) *VisualPropConverter {
	asc := NewArrowShapeConverter()
	sc := NewShapeConverter()

	valueConverterMap := map[string]ValueConverter {
		"arrow": asc,
		"shape": sc,
	}

	vpc := VisualPropConverter{typeTable:typeTable, valueConverters:valueConverterMap}

	return &vpc
}


func (vpConverter VisualPropConverter) GetCyjsPropertyValue(key string, value string) (converted interface{}) {

	dataType := vpConverter.typeTable[key]

	// Try converter
	converter, exists := vpConverter.valueConverters[dataType]
	if exists {
		return converter.Convert(value)
	}

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
