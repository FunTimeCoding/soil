package service_tester

import (
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"maps"
	"slices"
)

func Service(clients map[string]*mock_client.Client) *service.Service {
	var instances []instance.Instance
	names := slices.Sorted(maps.Keys(clients))

	for _, name := range names {
		instances = append(instances, instance.Instance{Name: name})
	}

	result := service.New(inventory.New[instance.Instance](instances...))

	for _, name := range names {
		result.Pool().Set(name, clients[name])
	}

	return result
}
