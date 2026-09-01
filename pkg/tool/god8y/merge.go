package god8y

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/text/dictionary"
)

func merge(sources []string) {
	target := dictionary.ResolvePath()
	added := dictionary.Merge(target, sources...)
	console.Format("Merged %d new words into %s\n", added, target)
}
