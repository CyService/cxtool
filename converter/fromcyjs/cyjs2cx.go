package fromcyjs

import (
	"bufio"
	"encoding/json"
	"os"
	"fmt"
	"github.com/cytoscape-ci/cxtool/cyjs"
	"io"
	"io/ioutil"
)


type Cyjs2Cx struct {
	W *io.Writer
}

func (con Cyjs2Cx) ConvertFromStdin() {
	if con.W == nil {
		outWriter := io.Writer(os.Stdout)
		con.W = &outWriter
	}

	reader := bufio.NewReader(os.Stdin)
	parseCx(reader, *con.W)
}

func (con Cyjs2Cx) Convert(sourceFileName string) {

	if con.W == nil {
		outWriter := io.Writer(os.Stdout)
		con.W = &outWriter
	}
	// Open the local file
	file, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Close input file at the end of this
	defer file.Close()

	reader := bufio.NewReader(file)

	parseCx(reader, *con.W)
}


func parseCx(reader *bufio.Reader, w io.Writer) {

	// Read everything into memory
	cyjsData, e := ioutil.ReadAll(reader)

	if e != nil {
		fmt.Printf("Read error: %v\n", e)
		return
	}

	var cytoscapeJsNetwork cyjs.CyJS

	err := json.Unmarshal(cyjsData, &cytoscapeJsNetwork)
	if err != nil {
		fmt.Println("error:", err)
	}

	cxNetwork := getCx(cytoscapeJsNetwork)

	cxJson, err := json.Marshal(cxNetwork)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	w.Write(cxJson)
}

