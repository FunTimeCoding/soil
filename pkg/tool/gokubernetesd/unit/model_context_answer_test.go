package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	kubernetesErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"testing"
)

func TestGettingAMissingPodAnswersWithoutAReport(t *testing.T) {
	c, r := service_tester.New(t).Mounted(true)
	text, e := c.CallToolError(
		constant.Get,
		map[string]any{
			"resourceType": "pod",
			"name":         "hotel",
			"namespace":    "default",
		},
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "hotel", text)
	assert.Count(t, 0, r.Events())
}

func TestAForbiddenDeleteAnswersWithoutAReport(t *testing.T) {
	v := service_tester.New(t)
	v.RefuseAll(
		kubernetesErrors.NewForbidden(
			schema.GroupResource{Resource: "pods"},
			"india",
			errors.New("rbac denied"),
		),
	)
	c, r := v.Mounted(false)
	text, e := c.CallToolError(
		constant.Delete,
		map[string]any{
			"resourceType": "pod",
			"name":         "india",
			"namespace":    "default",
		},
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "forbidden", text)
	assert.Count(t, 0, r.Events())
}

func TestAnInvalidPatchAnswersWithoutAReport(t *testing.T) {
	v := service_tester.New(t)
	v.RefuseAll(kubernetesErrors.NewBadRequest("the patch is not valid"))
	c, r := v.Mounted(false)
	text, e := c.CallToolError(
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
	assert.Count(t, 0, r.Events())
}

func TestAnUnexpectedClusterFailureIsReported(t *testing.T) {
	v := service_tester.New(t)
	v.RefuseAll(errors.New("etcd exploded"))
	c, r := v.Mounted(true)
	text, e := c.CallToolError(
		constant.Get,
		map[string]any{"resourceType": "pod", "namespace": "default"},
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "etcd exploded", text)
	assert.Count(t, 1, r.Events())
}
