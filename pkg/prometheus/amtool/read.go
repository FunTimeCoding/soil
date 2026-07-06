package amtool

import (
	"github.com/funtimecoding/soil/pkg/markup"
	"github.com/funtimecoding/soil/pkg/system"
)

func Read(
	base string,
	name string,
) *Option {
	c := &Option{}
	content := system.ReadFile(base, name)
	markup.MustDecode(content, &c)

	return c
}
