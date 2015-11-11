package converter

type Handler interface {

	Handle(item map[string]interface{})
}
