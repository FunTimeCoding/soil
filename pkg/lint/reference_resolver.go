package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
)

func ReferenceResolver(root string) *resolver.Resolver {
	o := option.New("", false)
	o.Metadata = true
	repository, _ := Walk(root, o)

	return newResolver(repository, o)
}
