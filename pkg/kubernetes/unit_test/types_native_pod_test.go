package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestPod(t *testing.T) {
	assert.NotNil(
		t,
		pod.New(
			&core.Pod{ObjectMeta: meta.ObjectMeta{Name: upper.Alfa}},
			"",
		),
	)
}
