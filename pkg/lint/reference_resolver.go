package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
)

func ReferenceResolver(root string) *pointer.Resolver {
	o := option.New("", false)
	o.Metadata = true
	repository, _ := Walk(root, o)

	return resolver(repository, o)
}
