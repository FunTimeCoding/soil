package mock_client

import (
	"encoding/base64"
	"github.com/funtimecoding/soil/pkg/gitlab/file"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) SeedFile(
	path string,
	content string,
) {
	c.files[path] = file.New(
		&gitlab.File{
			Content: base64.StdEncoding.EncodeToString([]byte(content)),
		},
	)
}
