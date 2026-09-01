package order

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/text/dictionary"
	"sort"
)

func Run() {
	f := dictionary.ResolvePath()
	c := dictionary.Read(f)

	for i := range c {
		sort.Strings(c[i].Words)
	}

	dictionary.Write(f, c)
	console.Format("Sorted %d categories in %s\n", len(c), f)
}
