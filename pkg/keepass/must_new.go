package keepass

import "github.com/funtimecoding/soil/pkg/errors"

func MustNew(
	database string,
	password string,
) *Client {
	result, e := New(database, password)
	errors.PanicOnError(e)

	return result
}
