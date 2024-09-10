package phone

import (
	"strings"
)

func FixMobilePhone(str string) string {
	str = strings.ReplaceAll(str, " ext. ", ",,")
	str = strings.ReplaceAll(str, " ", "")
	return str
}
