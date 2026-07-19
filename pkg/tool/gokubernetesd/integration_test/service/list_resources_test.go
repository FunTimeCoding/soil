package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester/pod"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"testing"
)

func TestListResources(t *testing.T) {
	s := service_tester.New(t)
	s.AddPod(pod.New("default", "nginx", "Running"))
	s.AddPod(pod.New("default", "redis", "Running").WithRestarts(5))
	result, e := s.Service.ListResources(
		context.Background(),
		"test",
		service.ListQuery{
			ResourceType: "pods",
			Namespace:    "default",
		},
	)
	assert.Nil(t, e)
	assert.Count(t, 2, result)
	assert.String(t, "nginx", result[0].Name)
	assert.String(t, "Running", result[0].Status)
}

func TestListResourcesWithRestarts(t *testing.T) {
	s := service_tester.New(t)
	s.AddPod(pod.New("default", "flaky", "Running").WithRestarts(47))
	result, e := s.Service.ListResources(
		context.Background(),
		"test",
		service.ListQuery{
			ResourceType: "pods",
			Namespace:    "default",
		},
	)
	assert.Nil(t, e)
	assert.Count(t, 1, result)
	assert.Integer(t, 47, int(result[0].Restarts))
}

func TestListResourcesEmpty(t *testing.T) {
	s := service_tester.New(t)
	result, e := s.Service.ListResources(
		context.Background(),
		"test",
		service.ListQuery{
			ResourceType: "pods",
			Namespace:    "default",
		},
	)
	assert.Nil(t, e)
	assert.Count(t, 0, result)
}

func TestListResourcesUnknownCluster(t *testing.T) {
	s := service_tester.New(t)
	_, e := s.Service.ListResources(
		context.Background(),
		"nonexistent",
		service.ListQuery{
			ResourceType: "pods",
			Namespace:    "default",
		},
	)
	assert.NotNil(t, e)
}
