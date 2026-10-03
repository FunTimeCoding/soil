package goc

import "github.com/funtimecoding/soil/pkg/console"

func SetEnvironmentEscape(
	k string,
	v string,
) {
	console.Format("\033]1337;SetUserVar=%s=%s\007", k, v)
}
