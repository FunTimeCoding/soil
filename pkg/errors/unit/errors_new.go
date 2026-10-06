package unit

import "errors"

func errorsNew(text string) error {
	return errors.New(text)
}
