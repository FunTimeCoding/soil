package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (t *Tester) AddNodeMetrics(
	name string,
	cpu string,
	memory string,
) {
	_, e := t.Dynamic.Resource(
		schema.GroupVersionResource{
			Group: "metrics.k8s.io", Version: "v1beta1", Resource: "nodes",
		},
	).Create(
		t.context(),
		&unstructured.Unstructured{
			Object: map[string]any{
				"apiVersion": "metrics.k8s.io/v1beta1",
				"kind":       "NodeMetrics",
				"metadata":   map[string]any{"name": name},
				"usage":      map[string]any{"cpu": cpu, "memory": memory},
			},
		},
		v1.CreateOptions{},
	)
	errors.PanicOnError(e)
}
