package converter

type SelectorEntry struct {
	
	Selector string `json:"selector"`
	CSS      map[string]interface{} `json:"css"`

}
