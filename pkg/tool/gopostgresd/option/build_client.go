package option

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/postgres"
	"github.com/funtimecoding/soil/pkg/postgres/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"net/url"
)

func BuildClient(i *instance.Instance) (face.PostgresClient, error) {
	result, e := postgres.New(
		fmt.Sprintf(
			"%s?%s",
			i.Locator(),
			url.Values{
				constant.ApplicationNameParameter: {constant.Identity.Name()},
			}.Encode(),
		),
	)

	if e != nil {
		return nil, fmt.Errorf("connect to %s: %w", i.Name, e)
	}

	return result, nil
}
