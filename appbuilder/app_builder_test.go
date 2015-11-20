package appbuilder

import (
	"testing"
	"reflect"
	"github.com/idekerlab/cxtool/converter"
)

func TestGetCoverter(t *testing.T) {

	typeCx2Cyjs := reflect.TypeOf(converter.Cx2Cyjs{})

	cv1 := getCoverter(cx)

	cv1Type := reflect.TypeOf(cv1)
	if cv1Type != typeCx2Cyjs {
		t.Errorf("Got: %q, expected %q", cv1Type, typeCx2Cyjs)
	}
}
