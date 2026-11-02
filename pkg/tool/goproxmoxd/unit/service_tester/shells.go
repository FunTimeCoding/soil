package service_tester

import "github.com/funtimecoding/soil/pkg/ssh/mock_client"

type Shells struct {
	given *mock_client.Client
	Built []*mock_client.Client
}
