package backup

import "context"

func plain(
	_ context.Context,
	e error,
) error {
	return e
}
