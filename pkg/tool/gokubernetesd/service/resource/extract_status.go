package resource

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ExtractStatus(u *unstructured.Unstructured) string {
	kind := u.GetKind()

	switch kind {
	case "Pod":
		return ExtractPodStatus(u)
	}

	status, _, e := unstructured.NestedMap(u.Object, "status")
	errors.PanicOnError(e)

	if status == nil {
		return ""
	}

	readyReplicas, okay, f := unstructured.NestedInt64(
		u.Object,
		"status",
		"readyReplicas",
	)
	errors.PanicOnError(f)

	if okay {
		replicas, _, g := unstructured.NestedInt64(
			u.Object,
			"status",
			"replicas",
		)
		errors.PanicOnError(g)

		return fmt.Sprintf("%d/%d ready", readyReplicas, replicas)
	}

	phase, okay, f := unstructured.NestedString(
		u.Object,
		"status",
		"phase",
	)
	errors.PanicOnError(f)

	if okay {
		return phase
	}

	conditions, okay, f := unstructured.NestedSlice(
		u.Object,
		"status",
		"conditions",
	)
	errors.PanicOnError(f)

	if okay {
		for _, raw := range conditions {
			condition, okay := raw.(map[string]any)

			if !okay {
				continue
			}

			if condition["type"] == "Ready" {
				if condition["status"] == "True" {
					return "Ready"
				}

				return "NotReady"
			}
		}
	}

	specType, okay, f := unstructured.NestedString(
		u.Object,
		"spec",
		"type",
	)
	errors.PanicOnError(f)

	if okay {
		return specType
	}

	return ""
}
