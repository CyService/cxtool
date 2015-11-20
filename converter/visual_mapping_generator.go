package converter
import (
	"log"
	"strings"
)

type VisualMappingGenerator struct {

}




func (vmGenerator VisualMappingGenerator) CreatePassthroughMapping(
vpName string, definition string, entry *SelectorEntry) {

	parts := strings.Split(definition, ",")
	values := strings.Split(parts[0], "=")


	log.Println(vpName, definition)
	if vpName == "NODE_LABEL" {
		entry.CSS[vpName] = "data(" + values[1] + ")"
	}
}
