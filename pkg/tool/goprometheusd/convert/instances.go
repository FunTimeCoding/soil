package convert

import "github.com/funtimecoding/soil/pkg/tool/goprometheusd/inventory"

func Instances(
	v []inventory.Instance,
	active string,
) []*SlimInstance {
	var result []*SlimInstance

	for i := range v {
		result = append(
			result,
			Instance(&v[i], v[i].Name == active),
		)
	}

	return result
}
