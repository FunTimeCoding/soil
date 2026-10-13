package inventory

import "github.com/funtimecoding/soil/pkg/inventory/face"

func (i *Inventory[T, P]) Validate() error {
	for k := range i.Instances {
		v, okay := any(P(&i.Instances[k])).(face.Validator)

		if !okay {
			continue
		}

		if e := v.Validate(); e != nil {
			return e
		}
	}

	return nil
}
