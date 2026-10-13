package mock_client

import (
	band "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
)

func (c *Client) EnableScreenService() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.screenService = band.ScreenServiceEnabled
	c.writes = append(c.writes, constant.ScreenServiceWrite)

	return nil
}
