package example

import "os/exec"

func OtherKey() {
	exec.Command("ls") // goanalyze:ignore restricted_call
}
