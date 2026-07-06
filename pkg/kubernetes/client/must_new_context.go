package client

import "github.com/funtimecoding/soil/pkg/errors"

func MustNewContext(cluster string) *Client {
	result, e := NewContext(cluster)
	errors.PanicOnError(e)

	return result
}
