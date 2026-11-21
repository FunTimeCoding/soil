package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"

func (c *Client) SeedRemediationItem(v client.RemediationItem) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.remediation = append(c.remediation, v)
}
