package layer

func New(
	color string,
	start float64,
	end float64,
) *Layer {
	return &Layer{Color: color, Start: start, End: end}
}
