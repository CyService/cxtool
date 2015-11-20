package converter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"bufio"
)

type Sif2Cx struct {
	Delimiter rune
}


func (con Sif2Cx) ConvertFromStdin() {
	log.Println("Waiting input from STDIN...")

	stdinReader := bufio.NewReader(os.Stdin)
	reader := csv.NewReader(stdinReader)
	con.readCSV(reader)
}

func (con Sif2Cx) Convert(sourceFileName string) {
	file, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Close input file at the end of this
	defer file.Close()
	reader := csv.NewReader(file)
	con.readCSV(reader)
}


func (con Sif2Cx) readCSV(reader *csv.Reader) {
	// Set delimiter
	reader.Comma = con.Delimiter
	reader.LazyQuotes = true

	// nodes already serialized
	nodesExists := make(map[string]int64)

	nodeCounter := int64(0)

	// The main CX document
//	cxDoc := CX{}

	fmt.Println("[")
	for {
		record, err := reader.Read()

		if err == io.EOF {
			netAttr := NetworkAttribute{N:"name", V:"SIF Conversion Test"}
			b, err := json.Marshal(netAttr)
			if err != nil {
				fmt.Println("error:", err)
			}

			fmt.Println("{\"networkAttributes\": [", string(b[:]), "]}]")

			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(record) == 3 {
			toJson(record, nodesExists, &nodeCounter)
		}
	}
}

func toJson(record []string, nodesExists map[string]int64, nodeCounter *int64) {

	source := record[0]
	interaction := record[1]
	target := record[2]


	_, exists := nodesExists[source]

	if !exists {
		*nodeCounter = *nodeCounter + 1
		nodesExists[source] = *nodeCounter
		sourceNode := Node{ID:*nodeCounter, N:source}
		printEntry(sourceNode)
	}

	_, targetExists := nodesExists[target]

	if !targetExists {
		*nodeCounter = *nodeCounter + 1
		nodesExists[target] = *nodeCounter
		targetNode := Node{ID:*nodeCounter, N:target}
		printEntry(targetNode)
	}

	*nodeCounter = *nodeCounter + 1
	edge := Edge{ID:*nodeCounter, S:nodesExists[source],
		T:nodesExists[target], I:interaction}


	printEdge(edge)
}

func printEdge(edge Edge) {
	newEdges := []Edge{edge}
	edgesEntry := Edges{EdgeList:newEdges}

	b, err := json.Marshal(edgesEntry)
	if err != nil {
		fmt.Println("error:", err)
	}

//	os.Stdout.Write(b)
	fmt.Println(string(b[:]), ",")
}

func printEntry(singleNode interface{}) {
	newNodes := []Node{singleNode.(Node)}
	nodesEntry := Nodes{NodesList:newNodes}

	b, err := json.Marshal(nodesEntry)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b[:]), ",")
}
