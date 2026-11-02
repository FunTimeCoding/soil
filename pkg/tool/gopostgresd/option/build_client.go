package option

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/postgres"
	"github.com/funtimecoding/soil/pkg/postgres/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
)

func BuildClient(i *instance.Instance) (face.Client, error) {
	result, e := postgres.New(i.Locator())

	if e != nil {
		return nil, fmt.Errorf("connect to %s: %w", i.Name, e)
	}

	return result, nil
}
