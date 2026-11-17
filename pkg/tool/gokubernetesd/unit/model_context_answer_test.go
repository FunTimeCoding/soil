package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	kubernetesErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"testing"
)

func TestGettingMissingPodAnswersWithoutReport(t *testing.T) {
	s := base.New(t)
	text, e := s.Tools().CallToolError(
		constant.Get,
		map[string]any{
			"resourceType": "pod",
			"name":         "hotel",
			"namespace":    "default",
		},
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "hotel", text)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestForbiddenDeleteAnswersWithoutReport(t *testing.T) {
	s := base.New(t, service_tester.WithArmed())
	s.RefuseAll(
		kubernetesErrors.NewForbidden(
			schema.GroupResource{Resource: "pods"},
			"india",
			errors.New("rbac denied"),
		),
	)
	text, e := s.Tools().CallToolError(
		constant.Delete,
		map[string]any{
			"resourceType": "pod",
			"name":         "india",
			"namespace":    "default",
		},
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "forbidden", text)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestInvalidPatchAnswersWithoutReport(t *testing.T) {
	s := base.New(t, service_tester.WithArmed())
	s.RefuseAll(kubernetesErrors.NewBadRequest("the patch is not valid"))
	text, e := s.Tools().CallToolError(
		constant.Patch,
		map[string]any{
			"resourceType": "pod",
			"name":         "juliett",
			"namespace":    "default",
			"patch":        "{}",
			"type":         "merge",
		},
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "not valid", text)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestUnexpectedClusterFailureIsReported(t *testing.T) {
	s := base.New(t)
	s.RefuseAll(errors.New("etcd exploded"))
	text, e := s.Tools().CallToolError(
		constant.Get,
		map[string]any{"resourceType": "pod", "namespace": "default"},
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "etcd exploded", text)
	assert.Count(t, 1, s.Reporter.Events())
}
