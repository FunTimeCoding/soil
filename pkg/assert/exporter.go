package assert

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
)

func exporter() cmp.Option {
	return cmp.Exporter(func(reflect.Type) bool { return true })
}
