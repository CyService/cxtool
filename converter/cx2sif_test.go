package converter_test


import (
	"testing"
	"bytes"
	"github.com/cytoscape-ci/cxtool/converter"
	"encoding/csv"
	"fmt"
	"strings"
)


//
// Read a CX file and generate SIF.
//
func TestCx2Sif(t *testing.T) {

	output := new(bytes.Buffer)
	csvWriter := csv.NewWriter(output)
	csvWriter.Comma = ' '

	c2s := converter.Cx2Sif{W:csvWriter}

	c2s.Convert("../test_data/galcxStyle2.json")

	result := output.String()

	lines := strings.Split(result, "\n")
	for idx := range lines {
		if lines[idx] == "" {
			fmt.Println(idx, ": EMPTY")
		} else {
			fmt.Println(idx, ": ", lines[idx])
		}
	}
	if len(lines) != 363 {
		t.Errorf("Expected 363, got %d", len(lines))
	}
}