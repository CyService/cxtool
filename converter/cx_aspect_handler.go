package converter

type CXAspectHandler interface {

	// Method to handle
	HandleAspect(aspect []interface{}) (attrMap map[string]interface{})
}
