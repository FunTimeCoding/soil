package answer

func New(
	line string,
	body string,
) *Answer {
	return &Answer{Line: line, Body: body}
}
