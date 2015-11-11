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

	file, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Close input file at the end of this
	defer file.Close()
	cxDecoder := json.NewDecoder(file)

	for {
		var val []map[string]interface{}
		err := cxDecoder.Decode(&val);

		if err == io.EOF {
			// Success!
			break
		} else if err != nil {
			log.Println(err)
			return
		}

		// Convert into Cytoscape.js
		cyjsNetwork := decodeCx(val)

		jsonString, err := json.Marshal(cyjsNetwork)
		if err != nil {
			fmt.Println("ERR: ", err)
		} else {
			//			fmt.Println("Cytoscpae.js JSON Generated:")
			fmt.Println(string(jsonString))
		}
	}
}


func decodeCx(val []map[string]interface{}) CyJS {

	// Network Object
	networkAttr := make(map[string]interface{})
	networkAttr["name"] = "Network Name"

	// Elements
	var nodes []CyJSNode
	var edges []CyJSEdge

	// Basic Cytoscape.js object
	cyjsNetwork := CyJS{Data: networkAttr}

	entryCount := len(val)

	// Iterate through the entire JSON
	for i := 0; i < entryCount; i++ {
		item := val[i]

		for key, value := range item {
			detectType(key, value, cyjsNetwork, &nodes, edges)
		}
	}

	elements := Elements{Nodes: nodes, Edges:edges}

	cyjsNetwork.Elements = elements

	log.Println("Last len = ", len(nodes))

	return cyjsNetwork
}


func detectType(tag string, value interface{}, cyjsNetwork CyJS,
cyjsNodes *[]CyJSNode, cyjsEdges []CyJSEdge) {

	switch tag {

	case networkAttributes:
		decodeNetworkAttributes(value.([]interface{}), cyjsNetwork)
	case nodes:
		decodeNodes(value.([]interface{}), cyjsNodes)
	case nodeAttributes:
		decodeNodeAttributes(value.([]interface{}))

	case edges:
		return
	case edgeAttributes:
	default:
	}
}

func decodeNetworkAttributes(value []interface{}, cyjsNetwork CyJS) {

	attrCount := len(value)
	for i := 0; i < attrCount; i++ {
		attr := value[i].(map[string]interface{})
		key := attr["n"].(string)
		cyjsNetwork.Data[key] = attr["v"]
	}

}

func decodeNodes(nodes []interface{}, cyjsNodes *[]CyJSNode) {

	nodeCount := len(nodes)

	for i := 0; i < nodeCount; i++ {
		node := nodes[i].(map[string]interface{})

		// Create data
		newNode := CyJSNode{}
		newNode.Data = make(map[string]interface{})
		newNode.Data["id"] = node[id].(string)
		newNode.Data["name"] = node[n].(string)

		*cyjsNodes = append(*cyjsNodes, newNode)

		jsonString, err := json.Marshal(newNode)

		if err != nil {
			fmt.Println("ERR: ", err)
		} else {
			log.Println(string(jsonString))
		}
	}
	log.Println("Cur LEN = ", len(*cyjsNodes))
}

func decodeNodeAttributes(attributes []interface{}) {

	attrCount := len(attributes)

	for i := 0; i < attrCount; i++ {
		//		attr := attributes[i].(map[string]interface{})

		//		jsonString, err := json.Marshal(attr)
		//		if err != nil {
		//			fmt.Println("ERR: ", err)
		//		} else {
		//			fmt.Println("Node Attr: ", string(jsonString))
		//		}
		//		fmt.Println("")
	}
}
