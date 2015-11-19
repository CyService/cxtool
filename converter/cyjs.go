package converter

// Basic Cytoscape.js network data structure.

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type CyJSNode struct {

	Data     map[string]interface{} `json:"data"`

	Position Position `json:"position,omitempty"`

	Selected bool `json:"selected,omitempty"`
}

type CyJSEdge struct {

	Data     map[string]interface{} `json:"data"`

	Selected bool `json:"selected,omitempty"`
}

type Elements struct {
	Nodes []CyJSNode `json:"nodes"`
	Edges []CyJSEdge `json:"edges"`
}


type CyJS struct {

	Data     map[string]interface{} `json:"data"`

	Elements Elements `json:"elements"`

	Style []SelectorEntry `json:"style"`
}