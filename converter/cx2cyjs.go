package converter

import (
	"os"
	"fmt"
	"encoding/json"
	"log"
	"io"
)

type Cx2Cyjs struct {
}

func (con Cx2Cyjs) Convert(sourceFileName, outputFileName string) {

	println("CX Source file: " + sourceFileName)
	println("Out File: " + outputFileName)

	file, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	cxDecoder := json.NewDecoder(file)

	for {
		var val []map[string]interface{}

		err := cxDecoder.Decode(&val);

		if err == io.EOF {
			fmt.Println("Success.")
			break
		} else if err != nil {
			log.Println("Error!!!!!!!!!")
			log.Println(err)
			return
		}
		fmt.Println("Length = ", len(val))
		decodeCx(val)

		//		fmt.Println(val)
	}
}


func decodeCx(val []map[string]interface{}) {

	entryCount := len(val)

	for i := 0; i < entryCount; i++ {

		item := val[i]


		for key, value := range item {
			detectType(key, value)
		}
	}
}

func detectType(tag string, value interface{}) {

	switch tag {

	case nodes:
		decodeNodes(value.([]interface{}))
	case nodeAttributes:
		decodeNodeAttributes(value.([]interface{}))
	case edges:
		fmt.Println(tag, ": ", value)
	default:
		fmt.Println(tag, ": Other. ", value)
	}

}


type Node struct {
    ID	string `json:"@id"`
    N string `json:"n"`
}

type NodeAttr struct {
	S string `json:"s"`
	Po string `json:"po"`
	N string `json:"n"`
	V string `json:"v"`
}

type Nodes struct {
	NODES []Node
}

/**
	Structure for Cytoscape.js Node
 */
type CyJSNode struct {
	Data map[string]interface{} `json:"data"`

	Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"position"`

	Selected bool `json:"selected"`
}


func decodeNodes(nodes []interface{}) {

	nodeCount := len(nodes)

	for i := 0; i < nodeCount; i++ {
		node := nodes[i].(map[string]interface{})

		// Create data
		data := make(map[string]interface{})
		data["id"] = node[id].(string)
		data["name"] = node[n].(string)

		newNode := CyJSNode{Data: data}

		fmt.Println("Node: ", newNode)
		jsonString, err := json.Marshal(newNode)
		if err != nil {
			fmt.Println("ERR: ", err)
		} else {
			fmt.Println("Node3: ", string(jsonString))
		}

	}
}

func decodeNodeAttributes(attributes []interface{}) {

	attrCount := len(attributes)

	for i := 0; i < attrCount; i++ {
		attr := attributes[i].(map[string]interface{})

		jsonString, err := json.Marshal(attr)
		if err != nil {
			fmt.Println("ERR: ", err)
		} else {
			fmt.Println("Node Attr: ", string(jsonString))
		}
		fmt.Println("")
	}
}
