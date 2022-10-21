package kata

import (
	"strconv"
	"strings"
)

func MultiTable(number int) string {
	filteredArr := []string{}
	for i := 1; i <= 10; i++ {
		filteredArr = append(filteredArr, strconv.Itoa(i)+" * "+strconv.Itoa(number)+" = "+strconv.Itoa(i*number))
	}

	return strings.Join(filteredArr, "\n")
}
