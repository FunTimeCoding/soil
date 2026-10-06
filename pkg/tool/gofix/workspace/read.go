package workspace

import "os"

func (w *Workspace) Read(path string) ([]byte, error) {
	if b, found := w.overlay[path]; found {
		return b, nil
	}

	return os.ReadFile(path)
}
