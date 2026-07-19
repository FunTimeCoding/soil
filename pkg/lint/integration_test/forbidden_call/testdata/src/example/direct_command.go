package example

import "os/exec"

func DirectCommand() {
	exec.Command("ls")
}
