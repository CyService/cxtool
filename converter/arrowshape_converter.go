package converter

type ArrowShapeConverter struct {
	arrowMap map[string]string
}

const (
	defaultArrowShape = "none"
)

func NewArrowShapeConverter() *ArrowShapeConverter {

	arrowMap := map[string]string{
		"T": "tee",
		"DELTA" : "triangle",
		"CIRCLE" : "circle",
		"DIAMOND": "diamond",
		"ARROW": "triangle",
		"HALF_BOTTOM": "triangle",
		"HALF_TOP": "triangle",
		"NONE": "none",
	}

	converter := ArrowShapeConverter{arrowMap:arrowMap}

	return &converter
}

func (conv ArrowShapeConverter) Convert(value string) string {

	cyjsArrowShape, exists := conv.arrowMap[value]

	if exists {
		return cyjsArrowShape
	} else {
		return defaultArrowShape
	}
}
