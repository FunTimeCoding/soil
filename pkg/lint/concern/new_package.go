package concern

import "github.com/funtimecoding/soil/pkg/lint/constant"

func NewPackage(
	key string,
	text string,
	path string,
) *Concern {
	return &Concern{
		Key:  key,
		Text: text,
		Path: path,
		Type: constant.ConcernPackage,
	}
}
