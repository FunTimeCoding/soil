package dial

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (d *Dial) background() string {
	base := d.track

	if d.filled && d.fillShare >= 1 {
		base = d.fillColor
	} else if d.filled && d.fillShare > 0 {
		base = fmt.Sprintf(
			"conic-gradient(%s 0deg %.1fdeg, %s %.1fdeg 360deg)",
			d.fillColor,
			degrees(d.fillShare),
			d.track,
			degrees(d.fillShare),
		)
	}

	parts := make([]string, 0, len(d.layers)+1)

	for _, l := range d.layers {
		parts = append(
			parts,
			fmt.Sprintf(
				"conic-gradient(transparent 0deg %.1fdeg, %s %.1fdeg %.1fdeg, transparent %.1fdeg 360deg)",
				degrees(l.from),
				l.color,
				degrees(l.from),
				degrees(l.to),
				degrees(l.to),
			),
		)
	}

	return join.CommaSpace(append(parts, base))
}
