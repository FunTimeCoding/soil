package dictionary

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"strings"
)

func Read(file string) []*Category {
	f := system.Open(file)
	defer errors.LogClose(f)
	var result []*Category
	var c *Category
	s := bufio.NewScanner(f)

	for s.Scan() {
		l := strings.TrimSpace(s.Text())

		if l == "" {
			continue
		}

		if strings.HasPrefix(l, "# ") {
			if c != nil {
				result = append(result, c)
			}

			c = &Category{Name: strings.TrimPrefix(l, "# ")}
		} else if c != nil {
			c.Words = append(c.Words, l)
		}
	}

	errors.PanicOnError(s.Err())

	if c != nil {
		result = append(result, c)
	}

	return result
}
