package example

import "os/exec"

func SuppressedSpace() {
	exec.Command("ls") // goanalyze:ignore
}
