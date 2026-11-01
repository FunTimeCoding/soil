package dial

import "maragu.dev/gomponents"

func (d *Dial) Center(n ...gomponents.Node) *Dial {
	d.center = append(d.center, n...)

	return d
}
