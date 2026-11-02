package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"

type Client struct {
	Result    *client.QueryResult
	Instances []client.Instance
	Failure   error
	Closed    int
}
