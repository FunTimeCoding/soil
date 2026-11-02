package file

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"

func New(c *commit.Service) *Service {
	return &Service{commit: c}
}
