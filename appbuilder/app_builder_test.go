package appbuilder

import (
	"testing"
	"reflect"
	"github.com/idekerlab/cxtool/converter"
	"bytes"
	"strings"
)

/*
	Main test cases to actually convert data
 */

func TestGetCoverter(t *testing.T) {

	typeCx2Cyjs := reflect.TypeOf(converter.Cx2Cyjs{})

	cv1 := getConverter(cx)

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
	app := BuildApp()
	testFileName := "../test_data/gal1.json"
	app.Run([]string{Name, testFileName})
}