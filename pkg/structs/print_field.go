package structs

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/structs/constant"
	"reflect"
)

func PrintField(f reflect.StructField) {
	console.Format(
		"name:%s type:%s tag:%s\n",
		f.Name,
		f.Type,
		f.Tag.Get(constant.NotationKey),
	)
}
