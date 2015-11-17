package converter

type CyjsStyle struct {

	Selectors []Section

}

type Section struct {
	
	Selector string
	CSS      map[string]interface{}
}
