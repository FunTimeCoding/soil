package label

import "github.com/funtimecoding/soil/pkg/errors/unexpected"

func statusFail(
	format string,
	status int,
) error {
	return unexpected.Format(format, status)
}
