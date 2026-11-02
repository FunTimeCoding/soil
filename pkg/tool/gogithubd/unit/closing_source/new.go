package closing_source

import "github.com/funtimecoding/soil/pkg/tool/gogithubd/mock_source"

func New() *Source {
	return &Source{Source: mock_source.New()}
}
