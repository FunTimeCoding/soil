package describe_result

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"

type Result struct {
	Resource map[string]any
	Events   []response.DescribeEvent
	Filtered []string
}
