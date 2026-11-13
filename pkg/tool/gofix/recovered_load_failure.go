package gofix

import "github.com/funtimecoding/soil/pkg/errors/load_failure"

func recoveredLoadFailure(v any) error {
	if v == nil {
		return nil
	}

	if e, okay := v.(error); okay {
		if failure, failed := load_failure.As(e); failed {
			return failure
		}
	}

	panic(v)
}
