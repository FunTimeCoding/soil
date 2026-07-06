package file_report

import "github.com/funtimecoding/soil/pkg/lint/concern"

func (r *Report) AddConcern(
	key string,
	text string,
	path string,
	line int,
	lineText string,
	fixed bool,
) {
	r.Add(concern.New(key, text, path, line, lineText, fixed))
}
