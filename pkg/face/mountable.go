package face

import "github.com/funtimecoding/soil/pkg/web/guard"

type Mountable interface {
	Mount(r Reporter, c Recorder, g *guard.Mux)
}
