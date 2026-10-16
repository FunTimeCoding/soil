package host_mapping

import "github.com/funtimecoding/soil/pkg/tool/goterraformd/types/plan"

func Addresses(p *plan.Plan) []string {
	var result []string

	for _, c := range p.ResourceChanges {
		if action(c.Change.Actions) != "" {
			result = append(result, c.Address)
		}
	}

	return result
}
