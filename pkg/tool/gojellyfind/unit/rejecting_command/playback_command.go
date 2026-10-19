package rejecting_command

import "github.com/funtimecoding/soil/pkg/errors/validation"

func (c *Client) PlaybackCommand(
	_ string,
	_ string,
	_ int64,
) error {
	return validation.New("One or more validation errors occurred.")
}
