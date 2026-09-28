package mock_lease_source

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"

type Client struct {
	leases []client.Lease
}
