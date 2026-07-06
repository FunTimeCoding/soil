package describe_result

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"

func New(
	resource map[string]any,
	events []response.DescribeEvent,
	filtered []string,
) *Result {
	return &Result{
		Resource: resource,
		Events:   events,
		Filtered: filtered,
	}
}
