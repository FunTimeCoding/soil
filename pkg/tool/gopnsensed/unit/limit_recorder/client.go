package limit_recorder

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/mock_source"

type Client struct {
	*mock_source.Client
	limit int
}
