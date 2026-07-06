package structs

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/structs/constant"
	"reflect"
)

func PrintField(f reflect.StructField) {
	fmt.Printf(
		"name:%s type:%s tag:%s\n",
		f.Name,
		f.Type,
		f.Tag.Get(constant.NotationKey),
	)
}
