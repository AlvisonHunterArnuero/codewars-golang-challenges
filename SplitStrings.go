package kata

import (
	"strings"
)

func Solution(str string) []string {
	pairedOne := strings.Split(str, "")
	result := make([]string, 0)

	if (len(pairedOne) % 2) != 0 {
		pairedOne = append(pairedOne, "_")
	}
	for i := 1; i < len(pairedOne); i += 2 {
		result = append(result, pairedOne[i-1]+pairedOne[i])
	}
	return result
}
