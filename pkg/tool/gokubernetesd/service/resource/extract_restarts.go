package resource

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ExtractRestarts(u *unstructured.Unstructured) int64 {
	containers, okay, e := unstructured.NestedSlice(
		u.Object,
		"status",
		"containerStatuses",
	)
	errors.PanicOnError(e)

	if !okay {
		return 0
	}

	var total int64

	for _, raw := range containers {
		c, okay := raw.(map[string]any)

		if !okay {
			continue
		}

		if count, okay := c["restartCount"].(int64); okay {
			total += count
		} else if count, okay := c["restartCount"].(float64); okay {
			total += int64(count)
		}
	}

	return total
}
