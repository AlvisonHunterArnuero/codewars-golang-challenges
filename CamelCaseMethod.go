package kata

import (
	"strings"
)

func CamelCase(s string) string {
	return strings.Join(strings.Split(strings.Title(strings.ToLower(s)), " "), "")
}
