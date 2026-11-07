package layout

import "github.com/funtimecoding/soil/pkg/identity"

func New(
	i *identity.Tool,
	theme string,
	favicon []byte,
) *Page {
	return &Page{identity: i, theme: theme, favicon: favicon}
}
