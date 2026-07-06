package relational_tag

import (
	"github.com/funtimecoding/soil/pkg/structs/constant"
	"slices"
)

func New(v []string) *Tag {
	defaultValue := ReadDefault(v)

	return &Tag{
		primary:      slices.Contains(v, constant.Primary),
		nullable:     !slices.Contains(v, constant.NotNull),
		hasDefault:   defaultValue != "",
		defaultValue: defaultValue,
	}
}
