package filter

import (
	"testing"
	"bytes"
	"io"
	"encoding/json"
	"strings"
	"os"
	"bufio"
)


func TestCx2IdList(t *testing.T) {

	output := new(bytes.Buffer)
	resultWriter := io.Writer(output)

	file, err := os.Open("../../test_data/galcxStyle2.json")
	if err != nil {
		t.Fatal("Error:", err)
		return
	}

	// Close input file at the end of this
	defer file.Close()

	cx2list := Cx2IdList{}
	cx2list.Convert(bufio.NewReader(file), resultWriter)
	result := output.String()

	var res map[string]int64
	err = json.NewDecoder(strings.NewReader(result)).Decode(&res)

	if err != nil {
		t.Fail()
	}

	t.Log(result)
	exp := 331

	actual := len(res)

	if actual == exp {
		t.Log("Pass")
	} else {
		t.Error("Expected: ", exp, "Actual: ", actual)
	}
}
