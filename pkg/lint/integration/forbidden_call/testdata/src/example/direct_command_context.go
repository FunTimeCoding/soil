package example

import (
	"context"
	"os/exec"
)

func DirectCommandContext() {
	exec.CommandContext(
		context.Background(),
		"ls",
	)
}
