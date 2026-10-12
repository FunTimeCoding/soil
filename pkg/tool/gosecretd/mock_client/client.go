package mock_client

import "github.com/funtimecoding/soil/pkg/openbao/status"

type Client struct {
	status    *status.Status
	health    error
	secrets   map[string]map[string]string
	failWrite error
}
