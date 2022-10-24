package kata

import (
	"strings"
)

func ModifyMultiply(str string, loc, num int) string {
	s := strings.Split(str, " ")
	newStr := s[loc]
	for i := 1; i < num; i++ {
		newStr += "" + "-" + s[loc]
	}
	return newStr
}
