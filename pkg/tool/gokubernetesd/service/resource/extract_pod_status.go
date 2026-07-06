package resource

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ExtractPodStatus(u *unstructured.Unstructured) string {
	reason, _, e := unstructured.NestedString(
		u.Object,
		"status",
		"phase",
	)
	errors.PanicOnError(e)
	r, okay, f := unstructured.NestedString(
		u.Object,
		"status",
		"reason",
	)
	errors.PanicOnError(f)

	if okay && r != "" {
		reason = r
	}

	containers, _, g := unstructured.NestedSlice(
		u.Object,
		"status",
		"containerStatuses",
	)
	errors.PanicOnError(g)

	for i := len(containers) - 1; i >= 0; i-- {
		c, okay := containers[i].(map[string]any)

		if !okay {
			continue
		}

		if waiting, okay := c["state"].(map[string]any)["waiting"]; okay {
			if w, okay := waiting.(map[string]any); okay {
				if r, okay := w["reason"].(string); okay && r != "" {
					reason = r
				}
			}
		} else if terminated, okay := c["state"].(map[string]any)["terminated"]; okay {
			if t, okay := terminated.(map[string]any); okay {
				if r, okay := t["reason"].(string); okay && r != "" {
					reason = r
				}
			}
		}
	}

	if u.GetDeletionTimestamp() != nil {
		reason = "Terminating"
	}

	return reason
}
