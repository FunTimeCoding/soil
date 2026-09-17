package repository

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func implicitBases(r *Repository) []string {
	var result []string

	if system.DirectoryExists(r.Absolute(constant.PackageDirectory)) {
		result = append(result, constant.PackageDirectory)
	}

	for _, sibling := range r.Siblings {
		p := filepath.Join(sibling, constant.PackageDirectory)

		if system.DirectoryExists(r.Absolute(p)) {
			result = append(result, p)
		}
	}

	return result
}
