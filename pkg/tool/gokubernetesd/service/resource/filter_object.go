package resource

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/format"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/resource/filter_result"
)

func FilterObject(
	object map[string]any,
	unfiltered bool,
) *filter_result.FilterResult {
	if unfiltered {
		return filter_result.New(object, nil)
	}

	var filtered []string
	metadata, okay := object["metadata"].(map[string]any)

	if !okay {
		return filter_result.New(object, nil)
	}

	if managed, okay := metadata["managedFields"]; okay {
		if entries, okay := managed.([]any); okay {
			filtered = append(
				filtered,
				fmt.Sprintf(
					"metadata.managedFields (%d entries)",
					len(entries),
				),
			)
		}

		delete(metadata, "managedFields")
	}

	annotations, okay := metadata["annotations"].(map[string]any)

	if okay {
		key := "kubectl.kubernetes.io/last-applied-configuration"

		if v, okay := annotations[key]; okay {
			size := 0

			if s, okay := v.(string); okay {
				size = len(s)
			} else {
				b, e := json.Marshal(v)

				if e == nil {
					size = len(b)
				}
			}

			filtered = append(
				filtered,
				fmt.Sprintf(
					"metadata.annotations[%s] (%s)",
					key,
					format.Bytes(size),
				),
			)
			delete(annotations, key)
		}
	}

	return filter_result.New(object, filtered)
}
