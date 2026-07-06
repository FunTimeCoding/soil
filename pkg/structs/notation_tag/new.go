package notation_tag

import (
	"github.com/funtimecoding/soil/pkg/structs/constant"
	"slices"
)

func New(v []string) *Tag {
	return &Tag{
		key:       ReadKey(v),
		omitEmpty: slices.Contains(v, constant.OmitEmpty),
	}
}
