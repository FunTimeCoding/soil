package closing_source

import "github.com/funtimecoding/soil/pkg/tool/gogithubd/mock_source"

type Source struct {
	*mock_source.Source
	closed int
}
