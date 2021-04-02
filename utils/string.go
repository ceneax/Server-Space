package utils

import (
	"regexp"
)

func Substring(str string, length int) string {
	reg := regexp.MustCompile(`!\[.*\]\(.+\)`)
	loc := reg.FindStringIndex(str)
	if loc != nil {
		substr := str[0:loc[0]]
		rs := []rune(substr)
		if length > len(rs) {
			return str[0:loc[1]]
		}
	}
	rs := []rune(str)
	rl := len(rs)
	if length > rl {
		return string(rs)
	} else {
		return string(rs[0:length])
	}

}
