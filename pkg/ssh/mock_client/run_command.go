package mock_client

import (
	"github.com/funtimecoding/soil/pkg/ssh/command"
	"github.com/funtimecoding/soil/pkg/system/result"
)

func (c *Client) RunCommand(o *command.Command) *result.Result {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Commands = append(c.Commands, o)

	return result.New("", "", 0, nil)
}
