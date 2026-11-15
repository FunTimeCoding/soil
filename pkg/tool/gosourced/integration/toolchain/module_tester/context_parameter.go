package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/addition"

func ContextParameter(
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
