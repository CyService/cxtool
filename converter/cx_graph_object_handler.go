package converter

type CXGraphObjectHandler interface {

	HandleGraphObject(aspect []interface{}) (graphObjects []interface{})

}
