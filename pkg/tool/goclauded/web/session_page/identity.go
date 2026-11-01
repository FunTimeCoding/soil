package session_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/web/component/fact"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) identity() gomponents.Node {
	d := p.view.Detail
	result := fact.New()

	if d.Session != nil && d.Session.Name != "" {
		result.Row("Name", d.Session.Name)
	}

	if d.Session != nil && d.Session.Callsign != nil {
		result.Row("Callsign", *d.Session.Callsign)
	}

	if d.Alias != "" {
		result.Row("Alias", d.Alias)
	}

	if d.Slug != "" {
		result.RowNode("Slug", html.Small(gomponents.Text(d.Slug)))
	}

	result.RowNode("Identifier", html.Small(gomponents.Text(d.Identifier)))

	if d.Created != "" {
		result.Row("Created", created(d.Created))
	}

	if d.TurnCount > 0 {
		result.Row("Turns", fmt.Sprintf("%d", d.TurnCount))
	}

	if d.Cost > 0 {
		result.Row("Cost", fmt.Sprintf("$%.2f", d.Cost))
	}

	if d.Session != nil {
		result.Row("Last Seen", time.Relative(d.Session.LastSeen))
	}

	return result.Node()
}
