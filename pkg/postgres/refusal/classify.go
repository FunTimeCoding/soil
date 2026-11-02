package refusal

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/postgres/constant"
	"github.com/jackc/pgx/v5/pgconn"
	"slices"
)

func Classify(e error) error {
	var p *pgconn.PgError

	if !errors.As(e, &p) || len(p.Code) < constant.ClassLength {
		return e
	}

	class := p.Code[:constant.ClassLength]

	switch {
	case slices.Contains(constant.UnavailableCodes, p.Code):
		return unavailable.Format("postgres: %s", p.Error())
	case p.Code == constant.QueryCanceledCode:
		return timeout.Format("postgres: %s", p.Error())
	case slices.Contains(constant.ForbiddenCodes, p.Code):
		return forbidden.Format("postgres: %s", p.Error())
	case slices.Contains(constant.ConflictClasses, class):
		return conflict.Format("postgres: %s", p.Error())
	case slices.Contains(constant.ValidationClasses, class):
		return validation.New("postgres: %s", p.Error())
	default:
		return e
	}
}
