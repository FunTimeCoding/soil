package embed

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"net/url"
)

func ModelEnvironment() string {
	if s := environment.Optional(constant.EmbedModelEnvironment); s != "" {
		return s
	}

	u, e := url.Parse(environment.Required(constant.EmbedTargetEnvironment))
	errors.PanicOnError(e)

	return DefaultModel(u.Scheme)
}
