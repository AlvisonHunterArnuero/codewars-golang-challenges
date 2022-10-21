package kata

func Feast(beast string, dish string) bool {
	var res bool

	if beast[0:1] == dish[0:1] && beast[len(beast)-1:] == dish[len(dish)-1:] {
		res = true
	}
	return res
}
