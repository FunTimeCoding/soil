package convert

import (
	"github.com/funtimecoding/soil/pkg/system/types/service"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/server"
)

func Services(v []*service.Service) []server.Service {
	result := make([]server.Service, 0, len(v))

	for _, s := range v {
		result = append(result, Service(s))
	}

	return result
}
