//go:build local

package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester/pod"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"testing"
)

func TestGetResource(t *testing.T) {
	s := service_tester.New(t)
	s.AddPod(pod.New("default", "nginx", "Running"))
	result, e := s.Service.GetResource(
		context.Background(),
		"test",
		service.GetQuery{
			ResourceType: "pods",
			Name:         "nginx",
			Namespace:    "default",
		},
	)
	assert.Nil(t, e)
	assert.NotNil(t, result.Object)
}

func TestGetResourceFiltered(t *testing.T) {
	s := service_tester.New(t)
	s.AddPod(pod.New("default", "nginx", "Running"))
	result, e := s.Service.GetResource(
		context.Background(),
		"test",
		service.GetQuery{
			ResourceType: "pods",
			Name:         "nginx",
			Namespace:    "default",
		},
	)
	assert.Nil(t, e)
	metadata, _ := result.Object["metadata"].(map[string]any)
	_, hasManagedFields := metadata["managedFields"]
	assert.False(t, hasManagedFields)
}

func TestGetResourceUnfiltered(t *testing.T) {
	s := service_tester.New(t)
	s.AddPod(pod.New("default", "nginx", "Running"))
	result, e := s.Service.GetResource(
		context.Background(),
		"test",
		service.GetQuery{
			ResourceType: "pods",
			Name:         "nginx",
			Namespace:    "default",
			Unfiltered:   true,
		},
	)
	assert.Nil(t, e)
	assert.NotNil(t, result.Object)
	assert.Count(t, 0, result.Filtered)
}

func TestGetResourceNotFound(t *testing.T) {
	s := service_tester.New(t)
	_, e := s.Service.GetResource(
		context.Background(),
		"test",
		service.GetQuery{
			ResourceType: "pods",
			Name:         "nonexistent",
			Namespace:    "default",
		},
	)
	assert.NotNil(t, e)
}
