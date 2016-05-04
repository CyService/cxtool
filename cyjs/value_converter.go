package cyjs

type ValueConverter interface {

	// Convert Cytoscape VP string value into Cyjs prop.
	Convert(value string) interface{}
}