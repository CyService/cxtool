package cx

import "strconv"

const (
	Cx_boolean = "boolean"
	Cx_byte = "byte"
	Cx_char = "char"
	Cx_double = "double"
	Cx_float = "float"
	Cx_integer = "integer"
	Cx_long = "long"
	Cx_short = "short"
	Cx_string = "string"
	Cx_listBoolean = "list_of_boolean"
	Cx_listByte = "list_of_byte"
	Cx_listChar = "list_of_char"
	Cx_listDouble = "list_of_double"
	Cx_listFloat = "list_of_float"
	Cx_listInteger = "list_of_integer"
	Cx_listLong = "list_of_long"
	Cx_listShort = "list_of_short"
	Cx_listString = "list_of_string"
)

type TypeDecoder struct {
}

func (decoder TypeDecoder) Decode(value string, dataType string) interface{} {

	switch dataType {
	case Cx_boolean:
		result, err := strconv.ParseBool(value)
		if err == nil {
			return result
		}
		return value
	case Cx_float, Cx_double:
		result, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return result
		}
		return value
	case Cx_integer, Cx_long, Cx_short:
		result, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return result
		}
		return value
	default:
		// If not supported, simply return string.
		// All lists will be handled as string.
		return value
	}
}
