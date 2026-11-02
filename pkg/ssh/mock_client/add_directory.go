package mock_client

func (c *Client) AddDirectory(path string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.directories[path] = true
}
