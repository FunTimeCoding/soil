package ssh

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/ssh/command"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func environmentPrefix(o *command.Command) string {
	var parts []string

	for k, v := range o.Environment {
		parts = append(parts, fmt.Sprintf("%s=%s;", k, v))
	}

	return join.Space(parts...)
}
