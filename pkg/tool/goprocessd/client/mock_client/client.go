package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"

type Client struct {
	ProcessList []client.Process
	Logs        map[string][]string
	Failure     error
	Reloads     int
	Closed      int
}
