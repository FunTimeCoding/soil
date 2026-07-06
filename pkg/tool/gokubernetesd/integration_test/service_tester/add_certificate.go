package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"time"
)

func (t *Tester) AddCertificate(
	namespace string,
	name string,
	dnsNames []string,
	notAfter time.Time,
	ready bool,
) {
	readyStatus := "False"

	if ready {
		readyStatus = "True"
	}

	var dnsNamesRaw []any

	for _, d := range dnsNames {
		dnsNamesRaw = append(dnsNamesRaw, d)
	}

	_, e := t.Dynamic.Resource(
		schema.GroupVersionResource{
			Group: "cert-manager.io", Version: "v1", Resource: "certificates",
		},
	).Namespace(namespace).Create(
		t.context(),
		&unstructured.Unstructured{
			Object: map[string]any{
				"apiVersion": "cert-manager.io/v1",
				"kind":       "Certificate",
				"metadata": map[string]any{
					"name":      name,
					"namespace": namespace,
				},
				"spec": map[string]any{"dnsNames": dnsNamesRaw},
				"status": map[string]any{
					"notAfter": notAfter.Format(time.RFC3339),
					"conditions": []any{
						map[string]any{
							"type":   "Ready",
							"status": readyStatus,
						},
					},
				},
			},
		},
		v1.CreateOptions{},
	)
	errors.PanicOnError(e)
}
