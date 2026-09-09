package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/credential"

func (c *Client) Search(_ string) []*credential.Credential {
	return nil
}
