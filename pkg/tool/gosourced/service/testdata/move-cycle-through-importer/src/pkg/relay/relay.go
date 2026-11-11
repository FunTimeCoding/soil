package relay

import "example/pkg/source"

func Relay() string {
	return source.Mode()
}
