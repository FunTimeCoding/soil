package service

import "golang.org/x/mod/modfile"

func requirements(b []byte) map[string]string {
	result := map[string]string{}
	f, e := modfile.ParseLax("", b, nil)

	if e != nil || f == nil {
		return result
	}

	for _, r := range f.Require {
		result[r.Mod.Path] = r.Mod.Version
	}

	for _, r := range f.Replace {
		if modfile.IsDirectoryPath(r.New.Path) {
			result[r.Old.Path] = ""
		}
	}

	return result
}
