package host_mapping

import "github.com/funtimecoding/soil/pkg/tool/goterraformd/types/plan"

func New(
	p *plan.Plan,
	attribute map[string]string,
) *Index {
	result := &Index{
		resources: map[string]*plan.Resource{},
		scope:     map[string]string{},
		calls:     map[string]*plan.Call{},
		parents:   map[string]string{},
		outputs:   map[string]map[string]*plan.Output{},
		values:    map[string][]map[string]any{},
		attribute: attribute,
	}

	if p.Configuration != nil && p.Configuration.RootModule != nil {
		result.addModule("", p.Configuration.RootModule)
	}

	if p.PriorState != nil && p.PriorState.Values != nil {
		result.addValues(p.PriorState.Values.RootModule)
	}

	if p.PlannedValues != nil {
		result.addValues(p.PlannedValues.RootModule)
	}

	return result
}
