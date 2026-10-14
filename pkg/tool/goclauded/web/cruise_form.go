package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strconv"
)

func (s *Server) cruiseForm(identifier string) gomponents.Node {
	setting, e := s.service.CruiseSetting(identifier)
	errors.PanicOnError(e)
	current := string(setting.Mode)

	if setting.Mode == constant.CruisePaced {
		current = strconv.Itoa(int(setting.Pace.Minutes()))
	}

	options := []gomponents.Node{
		cruiseOption(
			string(constant.CruiseOff),
			string(constant.CruiseOff),
			current,
		),
		cruiseOption(
			string(constant.CruiseOn),
			string(constant.CruiseOn),
			current,
		),
	}

	for _, p := range constant.CruisePace {
		minutes := strconv.Itoa(int(p.Minutes()))
		options = append(
			options,
			cruiseOption(
				minutes,
				fmt.Sprintf("paced %s minutes", minutes),
				current,
			),
		)
	}

	var woke gomponents.Node

	if setting.PromotedAt != nil {
		woke = html.Small(
			gomponents.Textf(
				"last woke %s",
				setting.PromotedAt.Local().Format("15:04"),
			),
		)
	}

	return html.Form(
		gomponents.Attr("method", "POST"),
		gomponents.Attr(
			"action",
			fmt.Sprintf("/sessions/%s/cruise", identifier),
		),
		html.Label(
			gomponents.Text("Cruise"),
			html.Select(html.Name(constant.Mode), gomponents.Group(options)),
		),
		html.Button(html.Type("submit"), gomponents.Text("Set")),
		woke,
	)
}
