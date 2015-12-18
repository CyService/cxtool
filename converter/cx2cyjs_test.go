package converter_test

import (
	"github.com/cytoscape-ci/cxtool/converter"
	cyjs "github.com/cytoscape-ci/cxtool/cyjs"
	"testing"
	"bytes"
	"io"
	"encoding/json"
	"strings"
	"errors"
	"strconv"
	"os"
	"bufio"
)


func TestCx2Cyjs(t *testing.T) {

	output := new(bytes.Buffer)
	resultWriter := io.Writer(output)

	file, err := os.Open("../test_data/galcxStyle2.json")
	if err != nil {
		t.Fatal("Error:", err)
		return
	}

	// Close input file at the end of this
	defer file.Close()

	c2c := converter.Cx2Cyjs{}
	c2c.Convert(bufio.NewReader(file), resultWriter)

	result := output.String()

	t.Log("Output length = ", len(result))

	pass := checkCytoscapeJSOutput(result, t)

	if pass {
		t.Log("Pass")
	} else {
		t.Error("Failed to validate Cytoscape.js output.")
	}
}

func checkCytoscapeJSOutput(serializedCyjsJSON string, t *testing.T) bool {

	// Decode in memory
	dec := json.NewDecoder(strings.NewReader(serializedCyjsJSON))
    var cyjsNetwork cyjs.CyJS
    dec.Decode(&cyjsNetwork)


	if cyjsNetwork.Data == nil {
		return false
	}
	if cyjsNetwork.CxData == nil {
		return false
	}
	if cyjsNetwork.Style == nil {
		return false
	}

	styleTest, err := testStyle(cyjsNetwork.Style)

	if err != nil {
		t.Log(err)
		return false
	} else if styleTest == false {
		return false
	}


	return true
}

func testStyle(style []cyjs.SelectorEntry) (bool, error) {

	if len(style) != 29 {
		errStr := "Number of Style entries is wrong: " + strconv.FormatInt(int64(len(style)), 10)
		return false, errors.New(errStr)
	}

	return true, nil
}

