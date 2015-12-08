/*
	Cytoscape.js Data model for serialization
 */

package cyjs


// Tags used in Cytoscape.js JSON
const (
	Id = "id"
	Source = "source"
	Target = "target"
	Interaction = "interaction"
)

// Position of a node
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Node Object
type CyJSNode struct {
	Data     map[string]interface{} `json:"data"`

	Position Position `json:"position,omitempty"`

	Selected bool `json:"selected"`
}


// Edge Object
type CyJSEdge struct {
	Data     map[string]interface{} `json:"data"`

	Selected bool `json:"selected,omitempty"`
}

// Elements in the network (nodes & edges)
type Elements struct {
	Nodes []CyJSNode `json:"nodes"`
	Edges []CyJSEdge `json:"edges"`
}

// Cytoscape.js Network Object
type CyJS struct {
	Data     map[string]interface{} `json:"data"`

	Elements Elements `json:"elements"`

	Style    []SelectorEntry `json:"style"`

	CxData   map[string]interface{} `json:"cxData,omitempty"`
}