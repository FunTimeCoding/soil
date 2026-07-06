package notation_tag

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/structs/constant"
	"reflect"
)

func Read(r reflect.StructField) []string {
	return split.Comma(r.Tag.Get(constant.NotationKey))
}
