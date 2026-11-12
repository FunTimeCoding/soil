package dial

import (
	"github.com/funtimecoding/soil/pkg/web/component/dial/layer"
	"github.com/funtimecoding/soil/pkg/web/component/dial/marker"
	"maragu.dev/gomponents"
)

type Dial struct {
	track     string
	fillColor string
	fillShare float64
	filled    bool
	layers    []*layer.Layer
	markers   []*marker.Marker
	center    []gomponents.Node
}
