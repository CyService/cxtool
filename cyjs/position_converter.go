package cyjs

import "strings"

type PositionConverter struct {
}

const (
	top = "top"
	center = "center"
	bottom = "bottom"
	left = "left"
	right = "right"

	NW = "NW"
	N = "N"
	NE = "NE"
	W = "W"
	C = "C"
	E = "E"
	SW = "SW"
	S = "S"
	SE = "SE"
)


func NewPositionConverter() *PositionConverter {

	converter := PositionConverter{}
	return &converter
}

func (conv PositionConverter) Convert(value string) interface{} {
	vals := strings.Split(value, ",")

	return getPos(vals[0])
}

func getPos(val string) [2]string {
	// First entry contains vertical, 2nd is horizontal.
	posPair := [2]string{}

	switch val {
	case N:
		posPair[0] = top
		posPair[1] = center
	case NW:
		posPair[0] = top
		posPair[1] = left
	case NE:
		posPair[0] = top
		posPair[1] = right
	case C:
		posPair[0] = center
		posPair[1] = center
	case W:
		posPair[0] = center
		posPair[1] = left
	case E:
		posPair[0] = center
		posPair[1] = right
	case S:
		posPair[0] = bottom
		posPair[1] = center
	case SW:
		posPair[0] = bottom
		posPair[1] = left
	case SE:
		posPair[0] = bottom
		posPair[1] = right
	default:
		posPair[0] = center
		posPair[1] = center
	}

	return posPair
}
