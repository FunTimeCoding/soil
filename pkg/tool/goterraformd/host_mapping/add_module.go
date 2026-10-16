package host_mapping

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/types/plan"
)

func (x *Index) addModule(
	prefix string,
	m *plan.Module,
) {
	x.outputs[prefix] = m.Outputs

	for _, r := range m.Resources {
		a := join.Empty(prefix, r.Address)
		x.resources[a] = r
		x.scope[a] = prefix
	}

	for name, c := range m.ModuleCalls {
		child := join.Empty(
			prefix,
			constant.ModuleScope,
			constant.ModuleSeparator,
			name,
			constant.ModuleSeparator,
		)
		x.calls[child] = c
		x.parents[child] = prefix

		if c.Module != nil {
			x.addModule(child, c.Module)
		}
	}
}
