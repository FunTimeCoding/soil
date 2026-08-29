package inventory

import "github.com/funtimecoding/soil/pkg/errors/validation"

func (i *Inventory) Validate() error {
	if len(i.Instances) < 2 {
		return nil
	}

	seen := map[int]string{}

	for _, v := range i.Instances {
		if name, okay := seen[v.Index]; okay {
			return validation.New(
				"instance index %d used by both %s and %s",
				v.Index,
				name,
				v.Name,
			)
		}

		seen[v.Index] = v.Name
	}

	return nil
}
