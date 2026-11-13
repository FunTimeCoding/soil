package response

type LinkResponse struct {
	Name string
}

type Note struct {
	Text string
}

func (n *Note) Body() string {
	return n.Text
}
