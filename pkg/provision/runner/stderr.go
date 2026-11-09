package runner

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/command"
)

func stderr(e error) string {
	if c, okay := errors.AsType[*command.CommandError](e); okay {
		return c.Stderr
	}

	return ""
}
