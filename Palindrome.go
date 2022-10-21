package kata

import (
	"strings"
)

func IsPalindrome(str string) bool {
	result := ""
	for _, v := range str {
		result = string(v) + result
	}
	return strings.ToUpper(str) == strings.ToUpper(result)
}
