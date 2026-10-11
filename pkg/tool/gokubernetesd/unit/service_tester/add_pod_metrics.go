package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (t *Tester) AddPodMetrics(
	namespace string,
	name string,
	containers []map[string]string,
) {
	var containerList []any

	for _, c := range containers {
		containerList = append(
			containerList,
			map[string]any{
				"name": c["name"],
				"usage": map[string]any{
					"cpu":    c["cpu"],
					"memory": c["memory"],
				},
			},
		)
	}

	_, e := t.Dynamic.Resource(
		schema.GroupVersionResource{
			Group:    "metrics.k8s.io",
			Version:  "v1beta1",
			Resource: "pods",
		},
	).Namespace(namespace).Create(
		t.context(),
		&unstructured.Unstructured{
			Object: map[string]any{
				"apiVersion": "metrics.k8s.io/v1beta1",
				"kind":       "PodMetrics",
				"metadata": map[string]any{
					"name":      name,
					"namespace": namespace,
				},
				"containers": containerList,
			},
		},
		v1.CreateOptions{},
	)
	errors.PanicOnError(e)
}
