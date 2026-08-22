package command

import "github.com/funtimecoding/soil/pkg/errors/constant"

func (e *CommandError) ErrorContext() (string, map[string]any) {
	return constant.Process, map[string]any{
		constant.Command: e.Command,
		constant.Output:  e.Output,
		constant.Stderr:  e.Stderr,
	}
}
