package dial

import "github.com/funtimecoding/soil/pkg/web/constant"

func (d *Dial) Tick(
	share float64,
	color string,
	title string,
) *Dial {
	d.markers = append(
		d.markers,
		&marker{
			share: share,
			class: constant.DialTickClass,
			color: color,
			title: title,
		},
	)

	return d
}
