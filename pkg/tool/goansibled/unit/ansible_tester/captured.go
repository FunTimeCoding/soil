package ansible_tester

import (
	"os"
	"path/filepath"
)

func (o *Tester) Captured(name string) (string, bool) {
	b, e := os.ReadFile(filepath.Join(o.capture, name))

	if e != nil {
		return "", false
	}

	return string(b), true
}
