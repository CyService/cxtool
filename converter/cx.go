package converter

const (
	nodes string = "nodes"
	edges string = "edges"

	networkAttributes string = "networkAttributes"
	nodeAttributes string = "nodeAttributes"
	edgeAttributes string = "edgeAttributes"
	cartesianLayout string = "cartesianLayout"

	// For nodes
	id string = "@id"
	n string = "n"
	s string = "s"
	t string = "t"
	i string = "i"
	po string = "po"
	v string = "v"
)

type Node struct {
	ID	string `json:"@id"`
	N string `json:"n"`
}

type NodeAttr struct {
	S string `json:"s"`
	Po string `json:"po"`
	N string `json:"n"`
	V string `json:"v"`
}
