package converter

import (
	cx "github.com/cyService/cxtool/cx"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"bufio"
)


type Sif2Cx struct {
	Delimiter rune
	Name string
}


func (con Sif2Cx) Convert(r io.Reader, w io.Writer) {
	reader := csv.NewReader(r)
	bufW := bufio.NewWriter(w)

	con.readSIF(reader, bufW)
}


func (con Sif2Cx) readSIF(reader *csv.Reader, w *bufio.Writer) {
	// Set delimiter

	var netName string
	if con.Name == "" {
		netName = "CX from SIF file"
	} else {
		netName = con.Name
	}

	reader.Comma = con.Delimiter
	reader.LazyQuotes = true

	// nodes already serialized
	nodesExists := make(map[string]int64)

	nodeCounter := int64(0)

	w.Write([]byte("["))

	for {
		record, err := reader.Read()

		if err == io.EOF {
			// Add network attributes at the end of doc.
			netAttr := cx.NetworkAttribute{N:"name", V: netName}

			attrList := []cx.NetworkAttribute{netAttr}
			netAttrs := make(map[string][]cx.NetworkAttribute)

			netAttrs["networkAttributes"] = attrList

			json.NewEncoder(w).Encode(netAttrs)

			w.Write([]byte("]"))
			w.Flush()
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if len(record) == 3 {
			toJson(record, nodesExists, &nodeCounter, w)
		}

		w.Flush()
	}
}


func toJson(record []string, nodesExists map[string]int64, nodeCounter *int64, w *bufio.Writer) {

	source := record[0]
	interaction := record[1]
	target := record[2]

	_, exists := nodesExists[source]

	if !exists {
		*nodeCounter = *nodeCounter + 1
		nodesExists[source] = *nodeCounter
		sourceNode := cx.Node{ID:*nodeCounter, N:source}
		printEntry(sourceNode, w)
	}

	_, targetExists := nodesExists[target]

	if !targetExists {
		*nodeCounter = *nodeCounter + 1
		nodesExists[target] = *nodeCounter
		targetNode := cx.Node{ID:*nodeCounter, N:target}
		printEntry(targetNode, w)
	}

	*nodeCounter = *nodeCounter + 1
	edge := cx.Edge{ID:*nodeCounter, S:nodesExists[source],
		T:nodesExists[target], I:interaction}

	printEdge(edge, w)
}


func printEdge(edge cx.Edge, w *bufio.Writer) {
	newEdges := []cx.Edge{edge}
	edgesEntry := cx.Edges{EdgeList:newEdges}
	json.NewEncoder(w).Encode(edgesEntry)
	w.WriteString(",")
}


func printEntry(singleNode interface{}, w *bufio.Writer) {
	newNodes := []cx.Node{singleNode.(cx.Node)}
	nodesEntry := cx.Nodes{NodesList:newNodes}
	json.NewEncoder(w).Encode(nodesEntry)
	w.WriteString(",")
}
