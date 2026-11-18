package face

import "context"

type Client interface {
	Reachable(x context.Context) error
	Batteries(x context.Context) ([]string, error)
	Variables(x context.Context, battery string) (map[string]string, error)
}
