package convert

import (
	"github.com/funtimecoding/soil/pkg/system/service"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/server"
)

func Service(s *service.Service) server.Service {
	return server.Service{
		Name:       s.Name,
		State:      s.State,
		Origin:     s.Origin,
		Source:     &s.Source,
		Version:    &s.Version,
		Deliberate: s.Deliberate,
	}
}
