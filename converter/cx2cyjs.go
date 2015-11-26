package converter

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime" // For debugging
	"strconv"
)

type ResourceReadError struct {
	message string
}

func (err ResourceReadError) Error() string {
	return err.message
}

type Cx2Cyjs struct {
}


func (con Cx2Cyjs) ConvertFromStdin() {
	reader := bufio.NewReader(os.Stdin)
	cxDecoder := json.NewDecoder(reader)
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
	run(cxDecoder)
}


func initHandlers() map[string]CXAspectHandler {

	table, typeTable, conversionErr := prepareConversionTable()

	if conversionErr != nil {
		return nil
	}

	handlers := make(map[string]CXAspectHandler)

	vpc := NewVisualPropConverter(typeTable)
	visualMappingGenerator := VisualMappingGenerator{vpConverter:*vpc}
	vpHandler := VisualStyleHandler{conversionTable: table,
		typeTable:typeTable, visualMappingGenerator:visualMappingGenerator}

	decoder := TypeDecoder{}
	attrHandler := AttributeHandler{typeDecoder:decoder}
	networkAttrHandler := NetworkAttributeHandler{typeDecoder: decoder}

	layoutHandler := LayoutHandler{}

	// Attribute Handlers: Use one common handler for all
	handlers[networkAttributes] = networkAttrHandler
	handlers[nodeAttributes] = attrHandler
	handlers[edgeAttributes] = attrHandler

	// Cytoscape specific handlers
	handlers[visualProperties] = vpHandler
	handlers[cartesianLayout] = layoutHandler

	return handlers
}

func prepareConversionTable() (conversionMap map[string]string, typeMap map[string]string, resourceErr error) {
	conversionTable, readErr := Asset("data/cx_to_cyjs_style.csv")
	if readErr != nil {
		return nil, nil, ResourceReadError{message: "Could not read resourcefile."}
	}

	table := make(map[string]string)
	typeTable := make(map[string]string)

	reader := csv.NewReader(bytes.NewReader(conversionTable))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(record) != 3 {
			continue
		} else {
			table[record[0]] = record[1]
			typeTable[record[0]] = record[2]
		}
	}

	return table, typeTable, nil
}

func run(cxDecoder *json.Decoder) {

	// Initialize handlers
	handlers := initHandlers()

	// Network Object
	networkAttr := make(map[string]interface{})

	// Elements
	var nodes []CyJSNode
	var edges []CyJSEdge
	layout := make(map[string]interface{})
	vps := make(map[string]interface{})

	elements := Elements{Nodes: nodes, Edges: edges}

	// Temp storage for attributes
	nodeAttrs := make(map[string]interface{})
	edgeAttrs := make(map[string]interface{})

	cxData := make(map[string]interface{})
	// Basic Cytoscape.js object
	cyjsNetwork := CyJS{Data: networkAttr, Elements: elements, CxData:cxData}

	for {
		_, err := cxDecoder.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			return
		}

		// Decode entry one-by-one.
		for cxDecoder.More() {
			var entry map[string]interface{}
			err := cxDecoder.Decode(&entry)
			if err != nil {
				log.Fatal(err)
			}

			parseCxEntry(handlers, entry, &cyjsNetwork, &nodeAttrs,
				&edgeAttrs, &layout, &vps)
		}
	}

	assignNodeAttr(cyjsNetwork.Elements.Nodes, nodeAttrs, layout)
	assignEdgeAttr(cyjsNetwork.Elements.Edges, edgeAttrs)

	// Add style to net
	cyjsNetwork.Style = vps["style"].([]SelectorEntry)

	jsonString, err := json.Marshal(cyjsNetwork)

	if err != nil {
		fmt.Println("ERR: ", err)
	} else {
		fmt.Println(string(jsonString))
	}
	//debug()
}

func parseCxEntry(
	handlers map[string]CXAspectHandler,
	entry map[string]interface{},
	cyjsNetwork *CyJS,
	nodeAttrs *map[string]interface{},
	edgeAttrs *map[string]interface{},
	layout *map[string]interface{},
	vps *map[string]interface{}) {

	for key, value := range entry {
		detectType(handlers, key, value, cyjsNetwork,
			nodeAttrs, edgeAttrs,
			layout, vps)
	}
}

func detectType(
	handlers map[string]CXAspectHandler,
	tag string, value interface{},
	cyjsNetwork *CyJS,
	nodeAttrs *map[string]interface{},
	edgeAttrs *map[string]interface{},
	layout *map[string]interface{},
	vps *map[string]interface{}) {

	switch tag {

	case networkAttributes:
		netHandler := handlers[networkAttributes]
		cyjsNetwork.Data = netHandler.HandleAspect(value.([]interface{}))
	case nodes:
		createNodes(value.([]interface{}), cyjsNetwork)
	case edges:
		decodeEdges(value.([]interface{}), cyjsNetwork)
	case nodeAttributes:
		nodeAttributeHandler := handlers[nodeAttributes]
		*nodeAttrs = nodeAttributeHandler.HandleAspect(value.([]interface{}))
	case edgeAttributes:
		edgeAttributeHandler := handlers[edgeAttributes]
		*edgeAttrs = edgeAttributeHandler.HandleAspect(value.([]interface{}))
	case cartesianLayout:
		layoutHandler := handlers[cartesianLayout]
		*layout = layoutHandler.HandleAspect(value.([]interface{}))
	case visualProperties:
		vpHandler := handlers[visualProperties]
		*vps = vpHandler.HandleAspect(value.([]interface{}))
	default:
		// All others
		cyjsNetwork.CxData[tag] = value
	}
}

func createNodes(nodes []interface{}, cyjsNetwork *CyJS) {
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
		newEdge.Data["source"] = strconv.FormatInt(int64(edge[s].(float64)), 10)
		newEdge.Data["target"] = strconv.FormatInt(int64(edge[t].(float64)), 10)

		itr, exists := edge[i]
		if exists {
			newEdge.Data["interaction"] = itr.(string)
		}

		*cyjsEdges = append(*cyjsEdges, newEdge)
	}
}

func assignNodeAttr(
	nodes []CyJSNode,
	nodeAttrs map[string]interface{}, layout map[string]interface{}) {

	nodeCount := len(nodes)

	for i := 0; i < nodeCount; i++ {
		nd := &nodes[i]
		nodeId := nd.Data["id"].(string)

		val, exists := nodeAttrs[nodeId]
		if exists {
			valueMap := val.(map[string]interface{})
			for key, value := range valueMap {
				nd.Data[key] = value
			}
		}

		// Assign position if available
		pos, exists := layout[nodeId]
		if exists {
			nd.Position = pos.(Position)
		}
	}
}

func assignEdgeAttr(
	edges []CyJSEdge,
	nodeAttrs map[string]interface{}) {

	edgeCount := len(edges)
	for i := 0; i < edgeCount; i++ {
		e := edges[i]
		nodeId := e.Data["id"].(string)

		val, exists := nodeAttrs[nodeId]
		if exists {
			valueMap := val.(map[string]interface{})
			for key, value := range valueMap {
				e.Data[key] = value
			}
		}
	}
}
func debug() {
	log.Println("--------- Memory Stats ------------")
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	log.Println(mem.Alloc/1000, " kb")
	log.Println(mem.TotalAlloc / 1000)
	log.Println(mem.HeapAlloc / 1000)
	log.Println(mem.HeapSys / 1000)
	log.Println("---------------------\n")
}
