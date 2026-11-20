package flag_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func DebianChips(r *overview_result.Row) gomponents.Node {
	if r.Unread != nil {
		return html.Small(gomponents.Text(r.Unread.Error()))
	}

	if len(r.Debian.Statuses) == 0 {
		return html.Small(gomponents.Text("no entries"))
	}

	chips := make([]gomponents.Node, 0, len(r.Debian.Statuses))

	for _, s := range r.Debian.Statuses {
		state := s.Status
		color := webConstant.BadColor
		title := s.Package

		switch {
		case s.NeverAffected:
			state = "not affected"
			color = webConstant.QuietColor
		case s.Fixed != "":
			state = "fixed"
			color = webConstant.OkayColor
			title = fmt.Sprintf("%s %s", s.Package, s.Fixed)
		}

		chips = append(
			chips,
			chip.New(fmt.Sprintf("%s %s", s.Release, state)).
				Color(color).
				Title(title).
				Node(),
		)
	}

	return gomponents.Group(chips)
}
