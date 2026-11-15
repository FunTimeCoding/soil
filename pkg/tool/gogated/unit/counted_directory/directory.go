package counted_directory

import "github.com/funtimecoding/soil/pkg/directory/mock_client"

type Directory struct {
	*mock_client.Client
	Authentications int
}
