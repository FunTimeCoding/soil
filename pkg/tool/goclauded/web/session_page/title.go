package session_page

import "github.com/funtimecoding/soil/pkg/strings/shorten"

func (p *Page) Title() string {
	d := p.view.Detail

	if d.Alias != "" {
		return d.Alias
	}

	if d.Slug != "" {
		return d.Slug
	}

	return shorten.Prefix(d.Identifier, 8)
}
