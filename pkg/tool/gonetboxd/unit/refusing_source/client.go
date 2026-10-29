package refusing_source

import "github.com/funtimecoding/soil/pkg/netbox/mock_client"

type Client struct {
	*mock_client.Client
	refusal error
}
