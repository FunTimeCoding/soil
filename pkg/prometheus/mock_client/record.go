package mock_client

import "time"

func (c *Client) record(moments ...time.Time) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.moments = append(c.moments, moments...)
}
