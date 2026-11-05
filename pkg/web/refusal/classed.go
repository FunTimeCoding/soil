package refusal

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func Classed(
	class string,
	message string,
) error {
	switch class {
	case errorsConstant.ValidationClass:
		return validation.New("%s", message)
	case errorsConstant.NotFoundClass:
		return not_found.Format("%s", message)
	case errorsConstant.ConflictClass:
		return conflict.Format("%s", message)
	case errorsConstant.ForbiddenClass:
		return forbidden.Format("%s", message)
	case errorsConstant.UnavailableClass:
		return unavailable.Format("%s", message)
	default:
		return unexpected.Format("%s", message)
	}
}
