package common

import "regexp"

var usernameRegex = regexp.MustCompile(`[a-zA-Z0-9_\-\.]*$`)

// Sanitize checks whether the struct is valid, returns a list of occured problems
func (u *User) Sanitize() []string {
	ret := make([]string, 0)
	ret = append(ret, SanitizeUsername(u.Username)...)
	return ret
}

func SanitizeUsername(s string) []string {
	ret := make([]string, 0)
	if len(s) < 6 {
		ret = append(ret, "Username has to be at least 6 characters long.")
	}
	if len(s) > 16 {
		ret = append(ret, "Username has to be at most 16 characters long.")
	}
	if !usernameRegex.MatchString(s) {
		ret = append(ret, "Username may only contain a-z, A-Z, 0-9, _, -, ..")
	}
	return ret
}

func SanitizeName(s string) []string {
	ret := make([]string, 0)
	if len(s) > 100 {
		ret = append(ret, "Name has to be at most 100 characters long.")
	}
	return ret
}
