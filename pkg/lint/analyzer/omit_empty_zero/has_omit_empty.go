package omit_empty_zero

import "strings"

func hasOmitEmpty(tag string) bool {
	options := strings.Split(tag, ",")

	for _, option := range options[1:] {
		if option == "omitempty" {
			return true
		}
	}

	return false
}
