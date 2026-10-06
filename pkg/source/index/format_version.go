package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index/xref"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"reflect"
	"sync"
)

var formatVersion = sync.OnceValue(
	func() string {
		return join.Space(
			constant.IndexFormatVersion,
			Shape(
				reflect.TypeFor[PackageRecord](),
				reflect.TypeFor[ExternalRecord](),
				reflect.TypeFor[xref.References](),
			),
		)
	},
)
