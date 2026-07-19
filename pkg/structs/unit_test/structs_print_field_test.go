package unit_test

import (
	"github.com/funtimecoding/soil/pkg/structs"
	"reflect"
	"testing"
)

func TestPrintField(t *testing.T) {
	type Test struct {
		Name string `json:"name"`
	}
	structs.PrintField(reflect.TypeOf(Test{}).Field(0))
}
