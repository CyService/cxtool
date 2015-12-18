package converter_test


import (
	"testing"
	"bytes"
	"github.com/cytoscape-ci/cxtool/converter"
	"fmt"
	"strings"
	"os"
	"bufio"
	"io"
)


//
// Read a CX file and generate SIF.
//
func TestCx2Sif(t *testing.T) {
	output := new(bytes.Buffer)
	w := io.Writer(output)

	file, err := os.Open("../test_data/galcxStyle2.json")
	if err != nil {
		t.Fatal("Error:", err)
		return
	}

	// Close input file at the end of this
	defer file.Close()

	c2s := converter.Cx2Sif{}
	c2s.Convert(bufio.NewReader(file), w)

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