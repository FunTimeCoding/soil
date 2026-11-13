package gocheck

import "golang.org/x/tools/go/packages"

func positioned(errors []packages.Error) packages.Error {
	for _, e := range errors {
		if e.Pos != "" {
			return e
		}
	}

	return errors[0]
}
