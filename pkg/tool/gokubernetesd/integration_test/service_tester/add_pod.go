package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester/pod"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"time"
)

func (t *Tester) AddPod(p *pod.Pod) {
	now := time.Now()
	var labelMap map[string]any

	if p.Labels != nil {
		labelMap = make(map[string]any)

		for k, v := range p.Labels {
			labelMap[k] = v
		}
	}

	metadata := map[string]any{
		"name":              p.Name,
		"namespace":         p.Namespace,
		"creationTimestamp": now.Format(time.RFC3339),
	}

	if labelMap != nil {
		metadata["labels"] = labelMap
	}

	_, e := t.Dynamic.Resource(
		schema.GroupVersionResource{
			Group: "", Version: "v1", Resource: "pods",
		},
	).Namespace(p.Namespace).Create(
		t.context(),
		&unstructured.Unstructured{
			Object: map[string]any{
				"apiVersion": "v1",
				"kind":       "Pod",
				"metadata":   metadata,
				"status": map[string]any{
					"phase": p.Phase,
					"containerStatuses": []any{
						map[string]any{
							"name":         p.Name,
							"restartCount": p.Restarts,
							"ready":        true,
							"state": map[string]any{
								"running": map[string]any{
									"startedAt": now.Format(time.RFC3339),
								},
							},
						},
					},
				},
			},
		},
		v1.CreateOptions{},
	)
	errors.PanicOnError(e)
	_, f := t.Clientset.CoreV1().Pods(p.Namespace).Create(
		t.context(),
		&core.Pod{
			ObjectMeta: v1.ObjectMeta{
				Name:      p.Name,
				Namespace: p.Namespace,
				Labels:    p.Labels,
			},
			Status: core.PodStatus{Phase: core.PodPhase(p.Phase)},
		},
		v1.CreateOptions{},
	)
	errors.PanicOnError(f)
}
