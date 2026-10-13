package service

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
)

func newEngine(i *instance.Instance) (face.Engine, error) {
	user := i.User

	if user == "" {
		user = constant.DefaultUser
	}

	return band.New(i.Locator(), user, i.Password), nil
}
