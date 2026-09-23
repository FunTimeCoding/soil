package target

import "example/pkg/gated"

func Describe() string {
	return gated.Value()
}

func Plain() string {
	return "plain"
}
