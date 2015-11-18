package converter

type CyjsStyle struct {

	Properties []SelectorEntry

}

type SelectorEntry struct {
	
	Selector string `json:"selector"`
	CSS      map[string]interface{} `json:"css"`

}
