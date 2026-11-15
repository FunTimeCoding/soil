package counted_directory

import "github.com/funtimecoding/soil/pkg/directory/mock_client"

func New() *Directory {
	return &Directory{Client: mock_client.New("")}
}
