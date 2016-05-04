package cyjs

type ShapeConverter struct {
	shapeMap map[string]string
}

const (
	defaultShape = "rectangle"
)


func NewShapeConverter() *ShapeConverter {
	shapeMap := map[string]string{
		"RECTANGLE": "rectangle",
		"ROUND_RECTANGLE": "roundrectangle",
		"TRIANGLE": "triangle",
		"PARALLELOGRAM": "rectangle",
		"DIAMOND": "diamond",
		"ELLIPSE": "ellipse",
		"HEXAGON": "hexagon",
		"OCTAGON": "octagon",
		"VEE":	"vee",
	}

	converter := ShapeConverter{shapeMap:shapeMap}

	return &converter
}


func (conv ShapeConverter) Convert(value string) interface{} {

	cyjsNodeShape, exists := conv.shapeMap[value]

	if exists {
		return cyjsNodeShape
	} else {
		return defaultShape
	}
}