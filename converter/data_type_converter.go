package converter

import "strconv"

const (
	cx_boolean     = "boolean"
	cx_byte        = "byte"
	cx_char        = "char"
	cx_double      = "double"
	cx_float       = "float"
	cx_integer     = "integer"
	cx_long        = "long"
	cx_short       = "short"
	cx_string      = "string"
	cx_listBoolean = "list_of_boolean"
	cx_listByte    = "list_of_byte"
	cx_listChar    = "list_of_char"
	cx_listDouble  = "list_of_double"
	cx_listFloat   = "list_of_float"
	cx_listInteger = "list_of_integer"
	cx_listLong    = "list_of_long"
	cx_listShort   = "list_of_short"
	cx_listString  = "list_of_string"
)

type TypeDecoder struct {
}

func (decoder TypeDecoder) decode(value string, dataType string) interface{} {

	switch dataType {
	case cx_boolean:
		result, err := strconv.ParseBool(value)
		if err == nil {
			return result
		}
		return value
	case cx_float, cx_double:
		result, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return result
		}
		return value
	case cx_integer, cx_long, cx_short:
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
