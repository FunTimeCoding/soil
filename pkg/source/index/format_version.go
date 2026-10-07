package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index/record"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"reflect"
	"sync"
)

var formatVersion = sync.OnceValue(
	func() string {
		return join.Space(
			constant.IndexFormatVersion,
			Shape(
				reflect.TypeFor[record.Package](),
				reflect.TypeFor[record.External](),
				reflect.TypeFor[record.References](),
			),
		)
	},
)
