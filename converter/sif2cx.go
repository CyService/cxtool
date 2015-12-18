package converter

import (
	cx "github.com/cytoscape-ci/cxtool/cx"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"bufio"
)


type Sif2Cx struct {
	Delimiter rune
}


func (con Sif2Cx) Convert(r io.Reader, w io.Writer) {
	reader := csv.NewReader(r)
	bufW := bufio.NewWriter(w)

	con.readSIF(reader, bufW)
}


func (con Sif2Cx) readSIF(reader *csv.Reader, w *bufio.Writer) {
	// Set delimiter

	reader.Comma = con.Delimiter
	reader.LazyQuotes = true

	// nodes already serialized
	nodesExists := make(map[string]int64)

	nodeCounter := int64(0)

	w.Write([]byte("["))

	for {
		record, err := reader.Read()

		if err == io.EOF {
			log.Println("-------- end ---------")
			// Add network attributes at the end of doc.
//			netAttr := cx.NetworkAttribute{N:"name", V:"SIF Conversion Test"}
//
//			attrList := []cx.NetworkAttribute{netAttr}
//			netAttrs := make(map[string][]cx.NetworkAttribute)
//
//			netAttrs["networkAttributes"] = attrList

			w.Write([]byte("]"))
			w.Flush()
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if len(record) == 3 {
			toJson(record, nodesExists, &nodeCounter, w)
		} else {
			log.Println("INVALID Line")
		}
	}

	log.Println("-------- end2 ---------")
}


func toJson(record []string, nodesExists map[string]int64, nodeCounter *int64, w *bufio.Writer) {

	log.Println(record)

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
	if len(newEdges) == 0 {
		log.Fatal("####################")
	}

	json.NewEncoder(w).Encode(edgesEntry)

	w.WriteString(",\n")
	w.Flush()
}


func printEntry(singleNode interface{}, w *bufio.Writer) {
	newNodes := []cx.Node{singleNode.(cx.Node)}
	nodesEntry := cx.Nodes{NodesList:newNodes}

	if len(newNodes) == 0 {
		log.Fatal("####################")
	}
	json.NewEncoder(w).Encode(nodesEntry)
	w.WriteString(",\n")
	w.Flush()
}
