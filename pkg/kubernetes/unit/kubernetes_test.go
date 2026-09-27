package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/markup"
	"github.com/funtimecoding/soil/pkg/kubernetes/notation"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/namespace"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/node"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	core "k8s.io/api/core/v1"
	events "k8s.io/api/events/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/rest"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(
		t,
		"KUBERNETES_CONTEXT",
		kubernetesConstant.ContextEnvironment,
	)
	assert.String(t, "type=Normal", kubernetesConstant.TypeNormal)
	assert.String(t, "pods", kubernetesConstant.PodsResource)
}

func TestValidateName(t *testing.T) {
	kubernetes.ValidateName(constant.LowerAlfa)
}

func TestFromConfiguration(t *testing.T) {
	c, e := client.FromConfiguration(
		&rest.Config{Host: "https://127.0.0.1:1"},
		"test",
	)
	assert.FatalOnError(t, e)
	assert.Nil(t, c.Validate())
	assert.String(t, "test", c.Cluster())
}

func TestRender(t *testing.T) {
	assert.Any(t, "null\n", markup.Render(nil))
}

func TestToUnstructured(t *testing.T) {
	assert.Any(
		t,
		&unstructured.Unstructured{Object: map[string]any{"kind": "Pod"}},
		notation.ToUnstructured([]byte(`{"kind":"Pod"}`)),
	)
}

func TestEvent(t *testing.T) {
	assert.NotNil(
		t,
		event.New(
			&events.Event{ObjectMeta: meta.ObjectMeta{Name: constant.UpperAlfa}},
			"",
		),
	)
}

func TestNamespace(t *testing.T) {
	assert.NotNil(
		t,
		namespace.New(
			&core.Namespace{
				ObjectMeta: meta.ObjectMeta{Name: constant.UpperAlfa},
			},
			"",
		),
	)
}

func TestNode(t *testing.T) {
	assert.NotNil(
		t,
		node.New(
			&core.Node{ObjectMeta: meta.ObjectMeta{Name: constant.UpperAlfa}},
			"",
		),
	)
}

func TestPod(t *testing.T) {
	assert.NotNil(
		t,
		pod.New(
			&core.Pod{ObjectMeta: meta.ObjectMeta{Name: constant.UpperAlfa}},
			"",
		),
	)
}
