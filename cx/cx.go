package cx

//
// Tags used in CX files
//
const (
// For nodes
	Id string = "@id"
	N  string = "n"
	S  string = "s"
	T  string = "t"
	I  string = "i"
	PO string = "po"
	V  string = "v"
)

//
// JSON structure for node
//
type Node struct {
	ID int64 `json:"@id"`
	N  string `json:"n"`
}

// List of Nodes
type Nodes struct {
	NodesList []Node `json:"nodes"`
}

// An Edge in CX
type Edge struct {
	ID int64 `json:"@id"`
	S  int64 `json:"s"`
	T  int64 `json:"t"`
	I  string `json:"i"`
}

// List of edges
type Edges struct {
	EdgeList []Edge `json:"edges"`
}


// Attribute for both nodes and edges
type Attribute struct {
	PO int64 `json:"po"`
	N  string `json:"n"`
	V  interface{} `json:"v"`
	D  string `json:"d,omitempty"`
}


// Special case: Attributes for a network
type NetworkAttribute struct {
	N string `json:"n"`
	V interface{} `json:"v"`
	D string `json:"d,omitempty"`
}

type Metadata struct {

}


// Basic structure of CX with minimum aspects for Cytoscape
type CX struct {

	NetworkAttributes []NetworkAttribute `json:"networkAttributes"`

	Nodes             Nodes `json:"nodes"`
	Edges             Edges `json:"edges"`

	NodeAttributes    []Attribute `json:"nodeAttributes"`
	EdgeAttributes    []Attribute `json:"edgeAttributes"`
}



//type NodeAttr struct {
//	S  string `json:"s"`
//	Po string `json:"po"`
//	N  string `json:"n"`
//	V  string `json:"v"`
//}