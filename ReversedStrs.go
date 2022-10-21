package kata

func Solution(word string) (reversedStr string) {
	for _, v := range word {
		reversedStr = string(v) + reversedStr
	}
	return
}
