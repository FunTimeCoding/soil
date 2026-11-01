package service

import (
	modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"golang.org/x/mod/modfile"
	"os"
	"path/filepath"
)

func modulePathOnDisk(directory string) string {
	b, e := os.ReadFile(filepath.Join(directory, modConstant.ModFile))

	if e != nil {
		return ""
	}

	return modfile.ModulePath(b)
}
