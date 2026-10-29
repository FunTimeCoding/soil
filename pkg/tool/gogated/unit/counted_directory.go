package unit

import "github.com/funtimecoding/soil/pkg/directory/mock_client"

type countedDirectory struct {
	*mock_client.Client
	authentications int
}
