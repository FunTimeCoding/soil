package dictionary

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/writer"
)

func Write(
	file string,
	v []*Category,
) {
	f := system.Create(file)
	defer errors.LogClose(f)
	r := bufio.NewWriter(f)
	defer errors.LogFlush(r)

	for i, c := range v {
		writer.Print(r, "# %s\n", c.Name)

		for _, w := range c.Words {
			writer.Print(r, "%s\n", w)
		}

		if i < len(v)-1 {
			writer.Print(r, "\n")
		}
	}
}
