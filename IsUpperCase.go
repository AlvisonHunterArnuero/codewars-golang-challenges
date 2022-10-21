package kata

import "regexp"

type MyString string

func (s MyString) IsUpperCase() bool {
	r, _ := regexp.Compile(`^[A-Z\s]+$`)
	return r.MatchString(string(s))
}
