package mock_outpost_source

import "github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"

type Client struct {
	host     client.Host
	services []client.Service
	fail     error
}
