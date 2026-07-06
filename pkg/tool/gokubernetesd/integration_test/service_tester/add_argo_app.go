package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (t *Tester) AddArgoApp(
	name string,
	namespace string,
	syncStatus string,
	healthStatus string,
) {
	_, e := t.Dynamic.Resource(
		schema.GroupVersionResource{
			Group:    "argoproj.io",
			Version:  "v1alpha1",
			Resource: "applications",
		},
	).Namespace(constant.Argocd).Create(
		t.context(),
		&unstructured.Unstructured{
			Object: map[string]any{
				"apiVersion": "argoproj.io/v1alpha1",
				"kind":       "Application",
				"metadata": map[string]any{
					"name":      name,
					"namespace": constant.Argocd,
				},
				"spec": map[string]any{
					"destination": map[string]any{"namespace": namespace},
				},
				"status": map[string]any{
					"sync":   map[string]any{"status": syncStatus},
					"health": map[string]any{"status": healthStatus},
				},
			},
		},
		v1.CreateOptions{},
	)
	errors.PanicOnError(e)
}
