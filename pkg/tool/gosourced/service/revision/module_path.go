package revision

import "golang.org/x/mod/modfile"

func modulePath(b []byte) string {
	return modfile.ModulePath(b)
}
