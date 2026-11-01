package log_page

import (
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
	"maragu.dev/gomponents"
)

func Listing(
	fights []fight.Fight,
	offset int,
	total int,
	start string,
	end string,
	paged bool,
) gomponents.Node {
	if !paged || len(fights) == 0 {
		return Table(fights)
	}

	return gomponents.Group(
		[]gomponents.Node{
			pagination(offset, total, start, end),
			Table(fights),
			pagination(offset, total, start, end),
		},
	)
}
