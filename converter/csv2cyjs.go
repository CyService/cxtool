package converter

import (
	"os"
	"fmt"
	"encoding/csv"
	"io"
	"log"
	"encoding/json"

)

type Csv2Cyjs struct {
	Delimiter rune
}

type Edge struct {
	Source, Target string
}

func (con Csv2Cyjs) Convert(sourceFileName string) {

	println("CSV Source file: " + sourceFileName)
	println(con.Delimiter)

	file, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	// Set delimiter
	reader.Comma = con.Delimiter
	reader.LazyQuotes = true

	lineCount := 0
	for {
		fmt.Println("Reading Line: ", lineCount)

		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

//		fmt.Println(record)

		toJson(record)

		lineCount += 1
	}
}




func toJson(record []string) {

	source := record[0]
	target := record[1]

	fmt.Println("Source = ", source)
	fmt.Println("Target = ", target)

	edge := Edge{Source: source, Target: target}

	b, err := json.Marshal(edge)
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)
}

