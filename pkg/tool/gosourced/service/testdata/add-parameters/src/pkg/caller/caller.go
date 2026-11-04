package caller

import (
	"context"
	"example/pkg/target"
)

var reply = target.Valued

func Handle(x context.Context) string {
	return target.Answer("a", 1)
}

func Wrapped(x context.Context) string {
	return target.Answer(
		"b",
		2,
	)
}

func Sum(x context.Context) int {
	return target.Grouped(1, 2)
}

func Background() string {
	return target.Unbound("c")
}

func Counted(x int) string {
	return target.Mistyped("d")
}

func Loose(value any) string {
	return target.Narrowed("n")
}

func Valuing(x context.Context) string {
	return reply("e") + target.Valued("f")
}

func Shadowed(x context.Context) string {
	return target.Shadow("g") + target.Log("h", "i") + target.Unnamed("j")
}

func Replying(x context.Context) string {
	var a target.Answerer = &target.Server{}

	return a.Reply("k") + (&target.Server{}).Reply("l")
}
