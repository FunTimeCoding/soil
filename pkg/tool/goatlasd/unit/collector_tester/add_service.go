package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
)

func AddService(
	f *mock_client.Client,
	name string,
	origin string,
	source string,
) {
	f.ServiceList = append(
		f.ServiceList,
		client.Service{
			Name:       name,
			State:      "running",
			Origin:     origin,
			Source:     &source,
			Deliberate: true,
		},
	)
}
