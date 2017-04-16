package shortener

import (
	"strconv"
	"strings"
)

func ConvertInt64ToString(i int64) string {

	return strings.ToUpper(strconv.FormatInt(i, 36))
}

func ConvertStringToInt64(s string) (int64, error) {
	return strconv.ParseInt(strings.ToLower(s), 36, 32)
}
