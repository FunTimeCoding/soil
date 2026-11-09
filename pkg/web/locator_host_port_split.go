package web

import "github.com/funtimecoding/soil/pkg/strings"

func LocatorHostPortSplit(s string) (string, int) {
	u := ParseLocator(s)

	return u.Hostname(), strings.MustToInteger(u.Port())
}
