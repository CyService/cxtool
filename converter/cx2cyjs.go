package converter

import (
	"os"
	"fmt"
	"encoding/json"
	"log"
	"io"
	"bufio"
	"runtime" // For debugging
	"strconv"
)

type Cx2Cyjs struct {
}

func (con Cx2Cyjs) ConvertFromStdin() {
	reader := bufio.NewReader(os.Stdin)
	cxDecoder := json.NewDecoder(reader)
	debug()
	run(cxDecoder)
}

func (con Cx2Cyjs) Convert(sourceFileName string) {

	file, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Close input file at the end of this
	defer file.Close()
	cxDecoder := json.NewDecoder(file)

	debug()

	run(cxDecoder)
}

func run(cxDecoder *json.Decoder) {

	// Network Object
	networkAttr := make(map[string]interface{})

	// Elements
	var nodes []CyJSNode
	var edges []CyJSEdge
	layout := make(map[string]Position)

	elements := Elements{Nodes:nodes, Edges:edges}

	// Temp storage for attributes
	nodeAttrs := make(map[string]map[string]interface{})
	edgeAttrs := make(map[string]map[string]interface{})

	// Basic Cytoscape.js object
	cyjsNetwork := CyJS{Data: networkAttr, Elements: elements}

	for {
		t, err := cxDecoder.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			return
		}

		log.Println("CX Array found: ", t)

		// Decode entry one-by-one.
		for cxDecoder.More() {
			var entry map[string]interface{}
			err := cxDecoder.Decode(&entry)
			if err != nil {
				log.Fatal(err)
			}

			parseCxEntry(entry, &cyjsNetwork, &nodeAttrs, &edgeAttrs, &layout)
		}
	}

	assignNodeAttr(cyjsNetwork.Elements.Nodes, nodeAttrs, layout)
	assignEdgeAttr(cyjsNetwork.Elements.Edges, edgeAttrs)

	jsonString, err := json.Marshal(cyjsNetwork)

	if err != nil {
		fmt.Println("ERR: ", err)
	} else {
		fmt.Println(string(jsonString))
	}
	debug()
}

func assignNodeAttr(nodes []CyJSNode,
nodeAttrs map[string]map[string]interface{}, layout map[string]Position) {
	nodeCount := len(nodes)
	for i := 0; i < nodeCount; i++ {
		nd := &nodes[i]
		nodeId := nd.Data["id"].(string)

		val, exists := nodeAttrs[nodeId]
		if exists {
			for key, value := range val {
				nd.Data[key] = value
			}
		}

		// Assign position if available
		pos, exists := layout[nodeId]
		if exists {
			nd.Position = pos
		}
	}
}

func assignEdgeAttr(edges []CyJSEdge,
nodeAttrs map[string]map[string]interface{}) {

	edgeCount := len(edges)
	for i := 0; i < edgeCount; i++ {
		e := edges[i]
		nodeId := e.Data["id"].(string)

		val, exists := nodeAttrs[nodeId]
		if exists {
			for key, value := range val {
				e.Data[key] = value
			}
		}
	}
}


func parseCxEntry(entry map[string]interface{},
cyjsNetwork *CyJS,
nodeAttrs *map[string]map[string]interface{},
edgeAttrs *map[string]map[string]interface{}, layout *map[string]Position) {

	for key, value := range entry {
		detectType(key, value, cyjsNetwork, nodeAttrs, edgeAttrs, layout)
	}
}



func detectType(tag string, value interface{},
cyjsNetwork *CyJS,
nodeAttrs *map[string]map[string]interface{},
edgeAttrs *map[string]map[string]interface{},
layout *map[string]Position) {

	switch tag {

	case networkAttributes:
//		decodeNetworkAttributes(value.([]interface{}), cyjsNetwork)
		netHandler := NetworkHandler{}
		netAttr := netHandler.HandleAspect(value.([]interface{}))
		cyjsNetwork.Data = netAttr
	case nodes:
		decodeNodes(value.([]interface{}), cyjsNetwork)
	case edges:
		decodeEdges(value.([]interface{}), cyjsNetwork)
	case nodeAttributes:
		decodeAttributes(value.([]interface{}), *nodeAttrs)
	case edgeAttributes:
		decodeAttributes(value.([]interface{}), *edgeAttrs)
	case cartesianLayout:
		decodeLayout(value.([]interface{}), *layout)
	default:
	}
}

func decodeLayout(entries []interface{}, layout map[string]Position) {
	layoutCount := len(entries)
	for i := 0; i < layoutCount; i++ {
		entry := entries[i].(map[string]interface{})
//		key := entry["node"].(string)
		key := strconv.FormatInt(int64(entry["node"].(float64)), 10)
		x := entry["x"].(float64)
		y := entry["y"].(float64)
		position := Position{X:x, Y:y}

		layout[key] = position
	}
}

func decodeNetworkAttributes(value []interface{}, cyjsNetwork *CyJS) {
	attrCount := len(value)
	for i := 0; i < attrCount; i++ {
		attr := value[i].(map[string]interface{})
		key := attr["n"].(string)
		cyjsNetwork.Data[key] = attr["v"]
	}
}

func decodeNodes(nodes []interface{}, cyjsNetwork *CyJS) {

	nodeCount := len(nodes)
	cyjsNodes := &cyjsNetwork.Elements.Nodes

	for i := 0; i < nodeCount; i++ {
		node := nodes[i].(map[string]interface{})

		// Create data
		newNode := CyJSNode{}
		newNode.Data = make(map[string]interface{})
		newNode.Data["id"] = strconv.FormatInt(int64(node[id].(float64)), 10)

		name, exists := newNode.Data["n"]
		if exists {
			newNode.Data["n"] = name.(string)
		}

		*cyjsNodes = append(*cyjsNodes, newNode)
	}
}

func decodeEdges(edges []interface{}, cyjsNetwork *CyJS) {

	edgeCount := len(edges)
	cyjsEdges := &cyjsNetwork.Elements.Edges

	for idx := 0; idx < edgeCount; idx++ {
		edge := edges[idx].(map[string]interface{})

		// Create data
		newEdge := CyJSEdge{}
		newEdge.Data = make(map[string]interface{})

		// Required fields
		newEdge.Data["id"] = strconv.FormatInt(int64(edge[id].(float64)), 10)
		newEdge.Data["source"] = strconv.FormatInt(int64(edge[s].(float64)),10)
		newEdge.Data["target"] = strconv.FormatInt(int64(edge[t].(float64)),10)

		itr, exists := edge[i]
		if exists {
			newEdge.Data["interaction"] = itr.(string)
		}

		*cyjsEdges = append(*cyjsEdges, newEdge)
	}
}

func decodeAttributes(attributes []interface{}, values map[string]map[string]interface{}) {

	attrCount := len(attributes)

	for i := 0; i < attrCount; i++ {
		attr := attributes[i].(map[string]interface{})

		// Extract pointer (key)
//		pointer := attr["po"].(string)
		pointer := strconv.FormatInt(int64(attr["po"].(float64)), 10)

		// Check the value already exists or not
		attrMap, exist := values[pointer]

		if !exist {
			attrMap = make(map[string]interface{})
		}

		attributeName := attr["n"].(string)
		attrMap[attributeName] = attr["v"]

		values[pointer] = attrMap
	}
}


func debug() {
	log.Println("--------- Memory Stats ------------")
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	log.Println(mem.Alloc / 1000, " kb")
	log.Println(mem.TotalAlloc / 1000)
	log.Println(mem.HeapAlloc / 1000)
	log.Println(mem.HeapSys / 1000)
	log.Println("---------------------\n")
}
