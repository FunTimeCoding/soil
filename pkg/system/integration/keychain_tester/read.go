package keychain_tester

import (
	"os"
	"path/filepath"
)

func (k *Tester) read(name string) string {
	b, e := os.ReadFile(filepath.Join(k.directory, name))

	if os.IsNotExist(e) {
		return ""
	}

	if e != nil {
		panic(e)
	}

	return string(b)
}
