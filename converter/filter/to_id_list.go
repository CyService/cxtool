package filter

import (
	cx "github.com/cytoscape-ci/cxtool/cx"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)


type Cx2IdList struct {
}


func (cx2l Cx2IdList) Convert(reader io.Reader, writer io.Writer) {
	cxDecoder := json.NewDecoder(reader)
	toList(cxDecoder, writer)
}


func toList(cxDecoder *json.Decoder, w io.Writer) {

	// Node ID to node name map
	nodeMap := make(map[string]int64)

	for {
		_, err := cxDecoder.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		for cxDecoder.More() {
			token, err := cxDecoder.Token()
			if err != nil {
				fmt.Println(err)
			}

			if token == "nodes" {
				processNode(cxDecoder, nodeMap)
			}
		}
	}

	err := json.NewEncoder(w).Encode(nodeMap)
	if err != nil {
		panic("parse err")
	}
}


func processNode(decoder *json.Decoder, nodes map[string]int64) {
	_, err := decoder.Token()
	if err != nil {
		return
	}

	var entry cx.Node

	for decoder.More() {
		err := decoder.Decode(&entry)
		if err != nil {
			return
		}

		if entry.N == "" {
			nodes[strconv.FormatInt(entry.ID, 10)] = entry.ID
		} else {
			nodes[entry.N] = entry.ID
		}
	}
}
