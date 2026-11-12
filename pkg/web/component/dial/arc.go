package dial

import "github.com/funtimecoding/soil/pkg/web/component/dial/layer"

func (d *Dial) Arc(
	color string,
	from float64,
	to float64,
) *Dial {
	d.layers = append(d.layers, layer.New(color, from, to))

	return d
}
