package service_tester

import "github.com/funtimecoding/soil/pkg/ssh/mock_client"

func newShells(given *mock_client.Client) *Shells {
	return &Shells{given: given}
}
