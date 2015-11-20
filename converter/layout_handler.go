package converter

import (
	"strconv"
)

type LayoutHandler struct {
}

func (layoutHandler LayoutHandler) HandleAspect(aspect []interface{}) map[string]interface{} {
	// Find length of this aspects to be processed
	layoutCount := len(aspect)

	layout := make(map[string]interface{})

	for i := 0; i < layoutCount; i++ {
		entry := aspect[i].(map[string]interface{})
		key := strconv.FormatInt(int64(entry[node].(float64)), 10)
		layout[key] = Position{X: entry["x"].(float64), Y: entry["y"].(float64)}
	}

	return layout
}
