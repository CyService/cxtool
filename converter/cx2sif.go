package converter
import (
	"bufio"
	"encoding/json"
	"os"
	"fmt"
	"io"
//	"reflect"
)

const (
	arrayStart = json.Delim('[')
)

type Cx2Sif struct {
}


func (con Cx2Sif) ConvertFromStdin() {
	reader := bufio.NewReader(os.Stdin)
	cxDecoder := json.NewDecoder(reader)
	writeSif(cxDecoder)
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
	writeSif(cxDecoder)
}


func writeSif(cxDecoder *json.Decoder) {

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
//			fmt.Println("Token = ", token)

			if token == "nodes" {
				processNode(cxDecoder)
			} else if token == "edges" {
				processEdge(cxDecoder)
			}
		}
	}
}

func processNode(decoder *json.Decoder) {
	token, err := decoder.Token()
	if err != nil || token != arrayStart {
//		fmt.Println("##INVALID: ", err)
		return
	}

	fmt.Println("Nodes found:---", token, "---")

	var entry interface{}
	for decoder.More() {
		err := decoder.Decode(&entry)
		if err != nil {
			fmt.Println("##ERR: ", err)
			return
		}

		fmt.Println("In Node --- ", entry)
	}

}

func processEdge(decoder *json.Decoder) {
	token, err := decoder.Token()
	if err != nil || token != arrayStart {
		return
	}

	var entry interface{}
	for decoder.More() {
		err := decoder.Decode(&entry)
		if err != nil {
			fmt.Println("##ERR: ", err)
			return
		}

		fmt.Println("In Edge --- ", entry)
	}

}
