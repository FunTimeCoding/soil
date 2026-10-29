package closing_client

import "github.com/funtimecoding/soil/pkg/tool/goitermd/mock_client"

type Client struct {
	*mock_client.Client
	closed int
}
