package backup

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/peered"
)

func connected(
	x context.Context,
	p *peered.Peered,
) error {
	sessions, e := p.Guarded(
		x,
		plain,
		constant.ConnectedSessionsQuery,
		fmt.Sprintf(constant.OwnApplicationPattern, constant.Identity.Name()),
	)

	if e != nil {
		return e
	}

	if len(sessions.Rows) == 0 {
		return nil
	}

	var described []string

	for _, row := range sessions.Rows {
		application := fmt.Sprint(row[0])

		if application == "" {
			application = constant.UnnamedApplication
		}

		described = append(
			described,
			fmt.Sprintf(
				constant.SessionDescription,
				application,
				row[1],
				row[2],
				row[3],
			),
		)
	}

	return conflict.Format(
		constant.SessionsConnected,
		join.CommaSpace(described),
	)
}
