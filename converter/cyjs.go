package converter

type CyJSNode struct {

	Data     map[string]interface{} `json:"data"`

	Position struct {
				 X float64 `json:"x"`
				 Y float64 `json:"y"`
			 } `json:"position"`

	Selected bool `json:"selected"`
}

type CyJSEdge struct {

	Data	map[string]interface{} `json:"data"`

	Selected bool `json:"selected"`
}

type Elements struct {
	Nodes []CyJSNode `json:"nodes"`
	Edges []CyJSEdge `json:"edges"`
}

type CyJS struct {
	Data     map[string]interface{} `json:"data"`

	Elements Elements `json:"elements"`
}
