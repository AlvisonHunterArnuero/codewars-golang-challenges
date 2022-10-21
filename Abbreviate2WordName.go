package kata

import (
	"strings"
)

func AbbrevName(name string) string {
	resultArr := strings.Fields(strings.ToUpper(name))
	return resultArr[0][0:1] + "." + resultArr[1][0:1]
}
