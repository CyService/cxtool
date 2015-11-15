package converter

type CXElementHandler interface {

	ProcessElement(element map[string]interface{})
}
