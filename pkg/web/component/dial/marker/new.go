package marker

func New(
	share float64,
	class string,
	color string,
	title string,
) *Marker {
	return &Marker{Share: share, Class: class, Color: color, Title: title}
}
