//go:build local

package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"k8s.io/apimachinery/pkg/types"
	"testing"
)

func TestPatchResource(t *testing.T) {
	s := service_tester.New(t)
	s.AddDeployment("default", "nginx", 1, 1)
	e := s.Service.PatchResource(
		context.Background(),
		"test",
		service.PatchQuery{
			ResourceType: "deployments",
			Name:         "nginx",
			Namespace:    "default",
			Patch:        `{"metadata":{"labels":{"patched":"true"}}}`,
			Type:         types.MergePatchType,
		},
	)
	assert.Nil(t, e)
	result, f := s.Service.GetResource(
		context.Background(),
		"test",
		service.GetQuery{
			ResourceType: "deployments",
			Name:         "nginx",
			Namespace:    "default",
			Unfiltered:   true,
		},
	)
	assert.Nil(t, f)
	metadata, _ := result.Object["metadata"].(map[string]any)
	labels, _ := metadata["labels"].(map[string]any)
	assert.String(t, "true", labels["patched"].(string))
}
