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

func TestDeleteResource(t *testing.T) {
	s := service_tester.New(t)
	s.AddPod(pod.New("default", "nginx", "Running"))
	e := s.Service.DeleteResource(
		context.Background(),
		"test",
		"pods",
		"nginx",
		"default",
	)
	assert.Nil(t, e)
	result, f := s.Service.ListResources(
		context.Background(),
		"test",
		service.ListQuery{
			ResourceType: "pods",
			Namespace:    "default",
		},
	)
	assert.Nil(t, f)
	assert.Count(t, 0, result)
}

func TestDeleteResourceNotFound(t *testing.T) {
	s := service_tester.New(t)
	e := s.Service.DeleteResource(
		context.Background(),
		"test",
		"pods",
		"nonexistent",
		"default",
	)
	assert.NotNil(t, e)
}
