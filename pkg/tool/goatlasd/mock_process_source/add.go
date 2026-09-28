package mock_process_source

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"

func (c *Client) Add(
	name string,
	running bool,
) {
	c.processes = append(
		c.processes,
		client.Process{Name: name, Command: name, Running: running},
	)
}
