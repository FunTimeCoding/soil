package backup

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/peered"
	"github.com/jackc/pgx/v5"
)

func (s *Service) beside(
	x context.Context,
	i *instance.Instance,
	p *peered.Peered,
	target string,
	path string,
) error {
	existing, e := p.Guarded(x, plain, constant.DatabaseExistsQuery, target)

	if e != nil {
		return e
	}

	if len(existing.Rows) > 0 {
		return conflict.Format(constant.DatabaseTaken, target)
	}

	if _, f := p.Guarded(
		x,
		plain,
		fmt.Sprintf(
			constant.CreateDatabaseStatement,
			pgx.Identifier{target}.Sanitize(),
		),
	); f != nil {
		return f
	}

	return s.dump.Restore(x, i, target, path, false)
}
