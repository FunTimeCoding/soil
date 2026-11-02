package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gobandd/constant"

func (c *Client) EnableListener() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.listener = true
	c.writes = append(c.writes, constant.ListenerWrite)

	return nil
}
