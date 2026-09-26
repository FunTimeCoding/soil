package service

import "path/filepath"

func resolve(path string) string {
	if v, e := filepath.EvalSymlinks(path); e == nil {
		return v
	}

	return path
}
