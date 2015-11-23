package converter_test


import (
	"testing"
	"github.com/idekerlab/cxtool/converter"
)

func TestCx2Sif(t *testing.T) {
	c2s := converter.Cx2Sif{}

	c2s.Convert("../test_data/ndex1.json")


}
