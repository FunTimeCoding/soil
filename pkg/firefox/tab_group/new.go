package tab_group

func New(
	title string,
	color string,
) *Group {
	return &Group{Title: title, Color: color}
}
