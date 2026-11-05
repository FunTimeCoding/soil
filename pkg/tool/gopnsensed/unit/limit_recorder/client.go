package limit_recorder

import "github.com/funtimecoding/soil/pkg/opnsense/mock_client"

type Client struct {
	*mock_client.Client
	limit int
}
