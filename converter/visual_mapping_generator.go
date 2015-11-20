package converter
import (
	"strings"
)

type VisualMappingGenerator struct {

}


func (vmGenerator VisualMappingGenerator) CreatePassthroughMapping(
vpName string, definition string, entry *SelectorEntry) {

	parts := strings.Split(definition, ",")
	if len(parts) != 2 {
		return
	}

	tagAndValue := strings.Split(parts[0], "=")
	if len(tagAndValue) != 2 {
		return
	}
	// This mapping is valid only for Labels (at least foe now...)
	if vpName == "NODE_LABEL" || vpName == "EDGE_LABEL" {
		entry.CSS[vpName] = "data(" + tagAndValue[1] + ")"
	}
}

