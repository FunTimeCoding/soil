package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	"testing"
)

func TestDescribeResource(t *testing.T) {
	s := service_tester.New(t)
	s.AddDeployment("default", "nginx", 1, 1)
	result, e := s.Service.DescribeResource(
		context.Background(),
		"test",
		request.Describe{
			ResourceType: "deployments",
			Name:         "nginx",
			Namespace:    "default",
		},
	)
	assert.Nil(t, e)
	assert.NotNil(t, result.Resource)
	assert.NotNil(t, result.Events)
}

func TestDescribeResourceFiltered(t *testing.T) {
	s := service_tester.New(t)
	s.AddDeployment("default", "nginx", 1, 1)
	result, e := s.Service.DescribeResource(
		context.Background(),
		"test",
		request.Describe{
			ResourceType: "deployments",
			Name:         "nginx",
			Namespace:    "default",
		},
	)
	assert.Nil(t, e)
	metadata, _ := result.Resource["metadata"].(map[string]any)
	_, hasManagedFields := metadata["managedFields"]
	assert.False(t, hasManagedFields)
}
