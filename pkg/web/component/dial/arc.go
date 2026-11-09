package dial

func (d *Dial) Arc(
	color string,
	from float64,
	to float64,
) *Dial {
	d.layers = append(d.layers, &layer{color: color, start: from, end: to})

	return d
}
