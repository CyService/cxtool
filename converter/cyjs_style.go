package converter

//
// An entry in Cytoscape.js style section
//
type SelectorEntry struct {

	Selector string                 `json:"selector"`
	CSS      map[string]interface{} `json:"css"`

}