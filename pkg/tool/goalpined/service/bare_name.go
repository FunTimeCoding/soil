package service

import "path/filepath"

func bareName(name string) bool {
	return filepath.IsLocal(name) && filepath.Base(name) == name
}
