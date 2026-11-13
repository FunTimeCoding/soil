package gobreak

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"golang.org/x/tools/cover"
)

func coverProfiles(path string) []*cover.Profile {
	if !system.FileExists(path) {
		return nil
	}

	result, e := cover.ParseProfiles(path)
	errors.PanicOnError(e)

	return result
}
