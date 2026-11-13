package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/source/index/kind"
)

func Kinds() []*kind.Kind {
	return []*kind.Kind{unclosed_resource.Kind(), face.Kind()}
}
