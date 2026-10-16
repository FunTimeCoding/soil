package host_mapping

import "github.com/funtimecoding/soil/pkg/tool/goterraformd/types/plan"

type Index struct {
	resources map[string]*plan.Resource
	scope     map[string]string
	calls     map[string]*plan.Call
	parents   map[string]string
	outputs   map[string]map[string]*plan.Output
	values    map[string][]map[string]any
	attribute map[string]string
}
