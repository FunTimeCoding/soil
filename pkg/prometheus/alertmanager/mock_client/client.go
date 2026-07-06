package mock_client

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"

type Client struct {
	alerts []*alert.Alert
}
