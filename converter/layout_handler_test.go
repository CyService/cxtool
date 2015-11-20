package converter

import (
	"testing"
)

func TestHandleAspect(t *testing.T) {

	var entry1 = map[string]interface{}{"node": float64(6048),
		"view": float64(5708), "x": float64(-735.8530296880053), "y": float64(1461.5375503050745)}
	var layout []interface{}

	layout = append(layout, entry1)

	layoutHandler := LayoutHandler{}

	result := layoutHandler.HandleAspect(layout)

	if len(result) != 1 {
		t.Errorf("Number of result error: ", len(result))
	}

	t.Log("Pass")
}
