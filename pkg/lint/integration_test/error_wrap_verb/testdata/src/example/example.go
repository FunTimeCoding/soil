package example

import (
	"errors"
	"fmt"
)

func Stringify(path string) error {
	e := errors.New("base")

	return fmt.Errorf("read %s: %v", path, e)
}

func StringVerb() error {
	e := errors.New("base")

	return fmt.Errorf("load: %s", e)
}

func DoubleWrap() error {
	e := errors.New("cause")

	return fmt.Errorf("%w: %v", errors.New("class"), e)
}

func Wrapped(path string) error {
	e := errors.New("base")

	return fmt.Errorf("read %s: %w", path, e)
}

func FlaggedVerb() error {
	e := errors.New("base")

	return fmt.Errorf("state: %+v", e)
}

func TypeVerb() error {
	e := errors.New("base")

	return fmt.Errorf("unexpected %T: %w", e, e)
}
