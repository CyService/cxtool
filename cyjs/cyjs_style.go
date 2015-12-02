/*
	An entry in Cytoscape.js Style object
 */
package cyjs

type SelectorEntry struct {

	Selector string                 `json:"selector"`
	CSS      map[string]interface{} `json:"css"`

}