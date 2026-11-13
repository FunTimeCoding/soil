package fact

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (f *Fact) glanceNode() gomponents.Node {
	var groups []gomponents.Node
	var title string
	var items []gomponents.Node
	grouped := false

	for _, r := range f.rows {
		if !r.Section {
			items = append(items, item(r))

			continue
		}

		if grouped || len(items) > 0 {
			groups = append(groups, f.group(title, items))
		}

		title = r.Label
		items = nil
		grouped = true
	}

	if !grouped {
		return f.grid(items)
	}

	groups = append(groups, f.group(title, items))

	return html.Div(groups...)
}
