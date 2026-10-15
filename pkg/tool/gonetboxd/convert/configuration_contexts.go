package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func ConfigurationContexts(
	v []*configuration_context.Context,
) []*server.ConfigurationContext {
	result := make([]*server.ConfigurationContext, 0, len(v))

	for _, c := range v {
		result = append(result, ConfigurationContext(c))
	}

	return result
}
