package converter
import (
	"bufio"
	"encoding/json"
	"os"
	"fmt"
	"io"
//	"reflect"
	"strconv"
	"encoding/csv"
)

const (
	arrayStart = json.Delim('[')
)

type Cx2Sif struct {
	W *csv.Writer
}

func (con Cx2Sif) ConvertFromStdin() {
	reader := bufio.NewReader(os.Stdin)
	cxDecoder := json.NewDecoder(reader)
	parseSif(cxDecoder, *con.W)
}

func (con Cx2Sif) Convert(sourceFileName string) {
	file, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Close input file at the end of this
	defer file.Close()
	cxDecoder := json.NewDecoder(file)
	parseSif(cxDecoder, *con.W)
}


func parseSif(cxDecoder *json.Decoder, w csv.Writer) {

	// Edge slice used for later mapping
	var edges []Edge

	// Node ID to node name map
	nodeMap := make(map[int64]string)

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
			} else if token == "edges" {
				processEdge(cxDecoder, &edges)
			}
		}
	}

	writeSif(nodeMap, edges, w)
}

func writeSif(nodes map[int64]string, edges []Edge, w csv.Writer) {
	for i := range edges {

		edge := edges[i]

		if edge.I == "" {
			w.Write([]string{nodes[edge.S], "i", nodes[edge.T]})
		} else {
			w.Write([]string{nodes[edge.S], edge.I, nodes[edge.T]})
		}
	}
	w.Flush()
}

func processNode(decoder *json.Decoder, nodes map[int64]string) {
	token, err := decoder.Token()
	if err != nil || token != arrayStart {
		return
	}

	var entry Node
	for decoder.More() {
		err := decoder.Decode(&entry)
		if err != nil {
			return
		}

		if entry.N == "" {
			nodes[entry.ID] = strconv.FormatInt(entry.ID, 10)
		} else {
			nodes[entry.ID] = entry.N
		}
	}

}

func processEdge(decoder *json.Decoder, edges *[]Edge) {
	token, err := decoder.Token()
	if err != nil || token != arrayStart {
		return
	}

	var entry Edge
	for decoder.More() {
		err := decoder.Decode(&entry)
		if err != nil {
			return
		}
		*edges = append(*edges, entry)
	}
}
