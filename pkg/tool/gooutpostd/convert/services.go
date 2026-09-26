package convert

import (
	"github.com/funtimecoding/soil/pkg/system/service"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/server"
)

func Services(
	v []*service.Service,
	origin string,
) []server.Service {
	result := make([]server.Service, 0, len(v))

	for _, s := range v {
		if origin != "" && s.Origin != origin {
			continue
		}

		result = append(result, Service(s))
	}

	return result
}
