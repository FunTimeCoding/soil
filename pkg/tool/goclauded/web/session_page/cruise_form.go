package session_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/component/choice"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strconv"
)

func (p *Page) CruiseForm(alert string) gomponents.Node {
	setting := p.view.Cruise
	current := string(setting.Mode)

	if setting.Mode == constant.CruisePaced {
		current = strconv.Itoa(int(setting.Pace.Minutes()))
	}

	choices := []*choice.Choice{
		choice.New(string(constant.CruiseOff), string(constant.CruiseOff)),
		choice.New(string(constant.CruiseOn), string(constant.CruiseOn)),
	}

	for _, pace := range constant.CruisePace {
		minutes := strconv.Itoa(int(pace.Minutes()))
		choices = append(
			choices,
			choice.New(minutes, fmt.Sprintf("paced %s minutes", minutes)),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			entry.New(
				ActionPath(p.view.Detail.Identifier, constant.CruiseAction),
				"Set",
			).
				Identifier(constant.CruiseAction).
				Inline().
				Alert(alert).
				Select(constant.Mode, "Cruise", choices, current).
				Node(),
			gomponents.If(
				setting.PromotedAt != nil,
				html.P(
					html.Class(webConstant.AbsentClass),
					gomponents.Textf(
						"woken %d times, last %s",
						setting.Woken,
						promoted(setting.PromotedAt),
					),
				),
			),
		},
	)
}
