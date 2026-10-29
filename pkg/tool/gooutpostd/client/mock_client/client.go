package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"

type Client struct {
	ServiceList []client.Service
	HostValue   *client.Host
	Failure     error
	Closed      int
}
