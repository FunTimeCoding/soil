package goc

import "github.com/funtimecoding/soil/pkg/console"

func GetEnvironmentEscape(k string) string {
	console.Format("\033]1337;GetUserVar=%s\007", k)

	return ""
}
