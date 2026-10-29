package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	kubernetesErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"testing"
)

func TestAnOverrideWhoseUpdateIsRefusedAnswersTheRefusal(t *testing.T) {
	v := service_tester.New(t)
	v.RefuseUpdate(
		kubernetesErrors.NewForbidden(
			schema.GroupResource{Resource: "configmaps"},
			"kilo",
			errors.New("rbac denied"),
		),
	)
	_, e := v.Service.ApplyResource(
		t.Context(),
		"test",
		request.Apply{
			Manifest: service_tester.OverrideManifest(),
			Override: true,
		},
	)
	assert.True(t, forbidden.Is(e))
	assert.Integer(t, 0, v.Creates())
}

func TestAnOverrideOfAMissingResourceCreatesIt(t *testing.T) {
	v := service_tester.New(t)
	_, e := v.Service.ApplyResource(
		t.Context(),
		"test",
		request.Apply{
			Manifest: service_tester.OverrideManifest(),
			Override: true,
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, v.Creates())
}
