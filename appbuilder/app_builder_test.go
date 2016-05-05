package appbuilder

import (
	"testing"
	"reflect"
	"github.com/cyService/cxtool/converter"
	"bytes"
	"strings"
)

/*
	Main test cases to actually convert data
 */

func TestGetConverter(t *testing.T) {

	typeCx2Cyjs := reflect.TypeOf(converter.Cx2Cyjs{})

	cv1 := getConverter(cx, cytoscapejs)

	cv1Type := reflect.TypeOf(cv1)
	if cv1Type != typeCx2Cyjs {
		t.Errorf("Got: %q, expected %q", cv1Type, typeCx2Cyjs)
	}
}

func TestHelp(t *testing.T) {
	app := BuildApp()

	// Build parameter
	output := new(bytes.Buffer)
	app.Writer = output
	app.Run([]string{Name})

	result := output.String()
	nameIndex := strings.IndexAny(result, "NAME")
	if nameIndex != 0 {
		t.Log("NANE Index: ", nameIndex)
		t.Errorf("unexpected output.  Should start with NAME: %s", result)
	}
}


func TestCx2Cyjs(t *testing.T) {
	output := new(bytes.Buffer)
	app := BuildApp()
	app.Writer = output
	testFileName := "../test_data/galcxStyle2.json"
	app.Run([]string{Name, testFileName})

	t.Log("\n", output)
	t.Log("\n--------OK------")
}


func TestCx2Sif(t *testing.T) {
	// Test conversion from CX to SIF file.

	output := new(bytes.Buffer)
	app := BuildApp()
	app.Writer = output
	testFileName := "../test_data/galcxStyle2.json"
	app.Run([]string{Name, "-o", "sif", testFileName})

	result := output.String()
//	vals := strings.Split(result, "\n")

	t.Log("Lines = ", len(result))
}
