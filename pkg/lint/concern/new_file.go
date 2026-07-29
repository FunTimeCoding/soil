package concern

import "github.com/funtimecoding/soil/pkg/lint/constant"

func NewFile(
	key string,
	text string,
	path string,
	fixed bool,
) *Concern {
	return &Concern{
		Key:   key,
		Text:  text,
		Path:  path,
		Type:  constant.ConcernFile,
		Fixed: fixed,
	}
}
