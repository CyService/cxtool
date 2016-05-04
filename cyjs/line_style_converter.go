package cyjs

type LineStyleConverter struct {
	lineStyleMap map[string]string
}

const (
	defaultLineStyle = "solid"
)


func NewLineStyleConverter() *LineStyleConverter {

	lineStyleMap := map[string]string{
		"SOLID": "solid",
		"DOT": "dotted",
		"DASH_DOT": "dotted",
		"LONG_DASH": "dashed",
		"EQUAL_DASH": "dashed",
	}

	converter := LineStyleConverter{lineStyleMap:lineStyleMap}
	return &converter
}


func (conv LineStyleConverter) Convert(value string) interface{} {

	cyjsLineStyle, exists := conv.lineStyleMap[value]

	if exists {
		return cyjsLineStyle
	} else {
		return defaultLineStyle
	}
}