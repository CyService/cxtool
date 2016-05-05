package fromcyjs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/cyService/cxtool/cyjs"
	"io"
	"io/ioutil"
)


type Cyjs2Cx struct {
}


func (con Cyjs2Cx) Convert(r io.Reader, w io.Writer) {

	bufR := bufio.NewReader(r)
	parseCx(bufR, w)
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

