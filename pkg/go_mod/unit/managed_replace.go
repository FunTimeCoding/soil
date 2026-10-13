package unit

import (
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
)

func managedReplace(
	path string,
	version string,
) *modfile.Replace {
	return &modfile.Replace{
		Old: module.Version{Path: path},
		New: module.Version{Path: path, Version: version},
	}
}
