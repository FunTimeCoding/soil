package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"testing"
)

func TestDescribeResource(t *testing.T) {
	s := service_tester.New(t)
	s.AddDeployment("default", "nginx", 1, 1)
	result, e := s.Service.DescribeResource(
		context.Background(),
		"test",
		service.DescribeQuery{
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
		service.DescribeQuery{
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
