package toolchain

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/addition"

func contextParameter(
	function string,
	receiver string,
	position int,
	argument string,
) *addition.Parameter {
	return addition.NewParameter(
		"example/pkg/target",
		function,
		receiver,
		"x",
		"context.Context",
		position,
		argument,
	)
}
