package dial

func (d *Dial) Fill(
	color string,
	share float64,
) *Dial {
	d.fillColor = color
	d.fillShare = share
	d.filled = true

	return d
}
