package goc

import "fmt"

func SetEnvironmentEscape(
	k string,
	v string,
) {
	// Not sure if this works
	fmt.Printf("\033]1337;SetUserVar=%s=%s\007", k, v)
}
