package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
)

type Client struct {
	identifier int
	alerts     []client.AlertsResponse
	failures   *failure.Failure
	Failure    error
	Closed     int
}
