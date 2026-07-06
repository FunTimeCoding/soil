package relational_tag

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/structs/constant"
	"reflect"
)

func Read(f reflect.StructField) []string {
	return split.Semicolon(f.Tag.Get(constant.RelationalKey))
}
