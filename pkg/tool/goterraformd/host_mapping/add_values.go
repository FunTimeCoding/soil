package host_mapping

import "github.com/funtimecoding/soil/pkg/tool/goterraformd/types/plan"

func (x *Index) addValues(m *plan.ValueModule) {
	if m == nil {
		return
	}

	for _, r := range m.Resources {
		a := configurationAddress(r.Address)
		x.values[a] = append(x.values[a], r.Values)
	}

	for _, c := range m.ChildModules {
		x.addValues(c)
	}
}
