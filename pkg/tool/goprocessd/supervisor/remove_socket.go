package supervisor

import "os"

func removeSocket(path string) {
	e := os.Remove(path)

	if e != nil {
		return
	}
}
