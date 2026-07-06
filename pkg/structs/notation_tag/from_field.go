package notation_tag

import (
	"github.com/funtimecoding/soil/pkg/structs/constant"
	"reflect"
)

func FromField(f reflect.StructField) *Tag {
	v := Read(f)

	if len(v) == 0 {
		return nil
	}

	if len(v) == 1 && v[0] == constant.Skip {
		return nil
	}

	return New(v)
}
