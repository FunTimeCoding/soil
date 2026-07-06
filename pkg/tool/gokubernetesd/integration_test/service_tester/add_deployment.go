package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	apps "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"time"
)

func (t *Tester) AddDeployment(
	namespace string,
	name string,
	readyReplicas int64,
	replicas int64,
) {
	r := int32(replicas)
	_, e := t.Dynamic.Resource(
		schema.GroupVersionResource{
			Group: "apps", Version: "v1", Resource: "deployments",
		},
	).Namespace(namespace).Create(
		t.context(),
		&unstructured.Unstructured{
			Object: map[string]any{
				"apiVersion": "apps/v1",
				"kind":       "Deployment",
				"metadata": map[string]any{
					"name":              name,
					"namespace":         namespace,
					"creationTimestamp": time.Now().Format(time.RFC3339),
				},
				"spec": map[string]any{
					"replicas": replicas,
					"selector": map[string]any{
						"matchLabels": map[string]any{
							"app": name,
						},
					},
				},
				"status": map[string]any{
					"readyReplicas": readyReplicas,
					"replicas":      replicas,
				},
			},
		},
		v1.CreateOptions{},
	)
	errors.PanicOnError(e)
	_, f := t.Clientset.AppsV1().Deployments(namespace).Create(
		t.context(),
		&apps.Deployment{
			ObjectMeta: v1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
			Spec: apps.DeploymentSpec{
				Replicas: &r,
				Selector: &v1.LabelSelector{
					MatchLabels: map[string]string{"app": name},
				},
			},
		},
		v1.CreateOptions{},
	)
	errors.PanicOnError(f)
}
