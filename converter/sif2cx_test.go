package converter

import (
	"testing"
	"bytes"
	"io"
	"os"
	"bufio"
	"encoding/json"
	"strings"
)


func TestSmall(t *testing.T) {
//	testSif2Cx("../test_data/galFiltered.sif", t)
	testSif2Cx("../test_data/small.sif", t)
}


func testSif2Cx(fName string, t *testing.T) {
	output := new(bytes.Buffer)
	w := io.Writer(output)

	f, err := os.Open(fName)

	if err != nil {
		t.Fatal("Error:", err)
		return
	}

	// Close input file at the end of this
	defer f.Close()

	s2c := Sif2Cx{Delimiter: rune(' ')}
	s2c.Convert(bufio.NewReader(f), w)
	result := output.String()

	t.Log(result)

	t.Log("Output length = ", len(result))

	dec := json.NewDecoder(strings.NewReader(result))
	var cx []interface{}
	dec.Decode(&cx)

	t.Log(cx)

}

