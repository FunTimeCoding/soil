package goc

import "github.com/funtimecoding/soil/pkg/console"

func GetEnvironmentEscape(k string) string {
	// Get, but returns empty?
	// Not sure how to then read it into a variable in go, if at all
	console.Format("\033]1337;GetUserVar=%s\007", k)

	return ""
}
