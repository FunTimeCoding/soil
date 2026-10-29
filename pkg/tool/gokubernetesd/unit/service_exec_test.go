package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	"testing"
)

func TestAnExecRefusesAnIncompleteRequest(t *testing.T) {
	v := service_tester.New(t)

	for _, q := range []request.Exec{
		{Namespace: "default", Command: []string{"true"}},
		{Name: "golf", Command: []string{"true"}},
		{Name: "golf", Namespace: "default"},
	} {
		_, e := v.Service.ExecInPod(t.Context(), "test", q)
		assert.True(t, validation.Is(e))
	}
}
