package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/file"

type Client struct {
	files   map[string]*file.File
	commits []*RecordedCommit
}
