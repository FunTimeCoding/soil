package instance_tester

import "github.com/funtimecoding/soil/pkg/errors/validation"

func (i *Instance) Validate() error {
	if i.Broken {
		return validation.New("instance %s is broken", i.Name)
	}

	return nil
}
