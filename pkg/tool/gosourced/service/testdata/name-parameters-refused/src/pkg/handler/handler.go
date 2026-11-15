package handler

import "context"

var x = "package"

func Declared(_ context.Context) string {
	x := "local"

	return x
}

func Outer(_ context.Context) string {
	return x
}

func Twice(_ context.Context, _ context.Context) {}

func Result(_ context.Context) (x string) {
	return ""
}
