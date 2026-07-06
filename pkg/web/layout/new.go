package layout

import "github.com/funtimecoding/soil/pkg/identity"

func New(i *identity.Tool) *Page {
	return &Page{identity: i}
}
