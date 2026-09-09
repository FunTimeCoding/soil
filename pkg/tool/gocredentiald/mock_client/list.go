package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/credential"

func (c *Client) List() []*credential.Credential {
	return nil
}
