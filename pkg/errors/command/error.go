package command

import "fmt"

func (e *CommandError) Error() string {
	return fmt.Sprintf("%s: %v", e.Command, e.Wrapped)
}
