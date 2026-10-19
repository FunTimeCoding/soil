package example

import (
	"example/pkg/errors/not_found"
	"example/pkg/errors/validation"
)

func isClientError(e error) bool {
	return not_found.Is(e) || validation.Is(e)
}

func isEmpty(e error) bool {
	return e == nil
}
