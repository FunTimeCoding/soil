package system

import "os"

func ExitOnCode(code int) {
	if code != 0 {
		os.Exit(code)
	}
}
