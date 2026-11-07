package option

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/web"
)

func (w *Web) Read(a *argument.Instance) {
	w.Address = a.Address()
	w.ServiceTokens = web.ServiceTokens()
}
