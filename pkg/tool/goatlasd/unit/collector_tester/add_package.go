package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
)

func AddPackage(
	f *mock_client.Client,
	name string,
	origin string,
	source string,
	packageName string,
	version string,
) {
	f.ServiceList = append(
		f.ServiceList,
		client.Service{
			Name:       name,
			State:      "running",
			Origin:     origin,
			Source:     &source,
			Package:    &packageName,
			Version:    &version,
			Deliberate: true,
		},
	)
}
