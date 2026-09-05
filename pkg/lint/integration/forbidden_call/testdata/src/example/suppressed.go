package example

import "os/exec"

func Suppressed() {
	exec.Command("ls") // goanalyze:ignore
}
