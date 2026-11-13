package lint

import "github.com/funtimecoding/soil/pkg/console"

func Header(
	name string,
	root string,
	detail string,
) {
	console.Line(HeaderLine(name, root, detail))
}
