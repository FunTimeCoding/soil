package mock_client

import (
	"github.com/funtimecoding/soil/pkg/band/constant"
	gobanddConstant "github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
)

func (c *Client) EnableScreenService() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.screenService = constant.ScreenServiceEnabled
	c.writes = append(c.writes, gobanddConstant.ScreenServiceWrite)

	return nil
}
