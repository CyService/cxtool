package converter

// Tags for CX

const (

// For nodes
	id string = "@id"
	n  string = "n"
	s  string = "s"
	t  string = "t"
	i  string = "i"
	po string = "po"
	v  string = "v"
)

type Node struct {
	ID int64 `json:"@id"`
	N  string `json:"n"`
}

type Nodes struct {
	NodesList []Node `json:"nodes"`
}

type Edge struct {
	ID int64 `json:"@id"`
	S  int64 `json:"s"`
	T  int64 `json:"t"`
	I  string `json:"i"`
}

type Edges struct {
	EdgeList []Edge `json:"edges"`
}

type Metadata struct {

}

type Attribute struct {
	PO int64
	N  string
	V  interface{}
	D  string
}

type NetworkAttribute struct {
	N string `json:"n"`
	V interface{} `json:"v"`
	D string `json:"d,omitempty"`
}


type CX struct {

	NetworkAttributes []NetworkAttribute `json:"networkAttributes"`

	Nodes             Nodes `json:"nodes"`
	Edges             Edges `json:"edges"`

	NodeAttributes    []Attribute `json:"nodeAttributes"`
	EdgeAttributes    []Attribute `json:"edgeAttributes"`

}



type NodeAttr struct {
	S  string `json:"s"`
	Po string `json:"po"`
	N  string `json:"n"`
	V  string `json:"v"`
}
