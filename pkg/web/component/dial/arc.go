package dial

func (d *Dial) Arc(
	color string,
	from float64,
	to float64,
) *Dial {
	d.layers = append(d.layers, &layer{color: color, from: from, to: to})

	return d
}
