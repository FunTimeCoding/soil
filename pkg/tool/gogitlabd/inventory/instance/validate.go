package instance

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/web"
)

func (i *Instance) Validate() error {
	if web.TrimScheme(i.Host) != constant.PublicHost {
		return nil
	}

	if len(i.Groups) > 0 || len(i.Projects) > 0 {
		return nil
	}

	return validation.New(
		"instance %s points at %s without groups or projects",
		i.Name,
		constant.PublicHost,
	)
}
