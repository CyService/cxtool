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
		fmt.Println(tag, ": ", value)
	case edges:
		fmt.Println(tag, ": ", value)
	default:
		fmt.Println(tag, ": Other. ", value)
	}

}



