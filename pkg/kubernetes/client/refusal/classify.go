package refusal

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"k8s.io/apimachinery/pkg/api/errors"
)

func Classify(e error) error {
	if e == nil {
		return nil
	}

	if f := connection.Classify(e); f != nil {
		return f
	}

	switch {
	case errors.IsNotFound(e):
		return not_found.Format("%s", e.Error())
	case errors.IsForbidden(e), errors.IsUnauthorized(e):
		return forbidden.Format("%s", e.Error())
	case errors.IsConflict(e), errors.IsAlreadyExists(e):
		return conflict.Format("%s", e.Error())
	case errors.IsInvalid(e), errors.IsBadRequest(e):
		return validation.New("%s", e.Error())
	case errors.IsServiceUnavailable(e):
		return unavailable.Format("%s", e.Error())
	default:
		return e
	}
}
