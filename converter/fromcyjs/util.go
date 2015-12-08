package fromcyjs

import "strconv"

func getIdNumber(idStr string) (int64, error) {
	return strconv.ParseInt(idStr, 10, 64)
}
