package god8y

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/text/dictionary"
)

func merge(sources []string) {
	target := dictionary.ResolvePath()
	added := dictionary.Merge(target, sources...)
	fmt.Printf("Merged %d new words into %s\n", added, target)
}
