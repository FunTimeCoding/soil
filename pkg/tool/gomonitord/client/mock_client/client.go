package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"

type Client struct {
	ClaimList []client.Claim
	Failure   error
	Closed    int
}
