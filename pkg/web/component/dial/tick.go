package dial

import (
	"github.com/funtimecoding/soil/pkg/web/component/dial/marker"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (d *Dial) Tick(
	share float64,
	color string,
	title string,
) *Dial {
	d.markers = append(
		d.markers,
		marker.New(share, constant.DialTickClass, color, title))

	return d
}
