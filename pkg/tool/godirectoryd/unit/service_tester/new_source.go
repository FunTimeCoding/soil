package service_tester

import "github.com/funtimecoding/soil/pkg/directory/mock_client"

func NewSource() *mock_client.Client {
	return mock_client.New("dc=example,dc=org")
}
