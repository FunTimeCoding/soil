package concern

import "github.com/funtimecoding/soil/pkg/lint/constant"

func NewLine(
	key string,
	text string,
	path string,
	line int,
	lineText string,
	fixed bool,
) *Concern {
	return &Concern{
		Key:      key,
		Text:     text,
		Path:     path,
		Type:     constant.ConcernLine,
		Line:     line,
		LineText: lineText,
		Fixed:    fixed,
	}
}
