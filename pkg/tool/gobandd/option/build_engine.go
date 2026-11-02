package option

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
)

func BuildEngine(i *instance.Instance) (face.BandClient, error) {
	if i.Host == "" || i.Password == "" {
		return nil, validation.New(
			"band instance %s needs a host and a password",
			i.Name,
		)
	}

	user := i.User

	if user == "" {
		user = constant.DefaultUser
	}

	return band.New(i.Locator(), user, i.Password), nil
}
