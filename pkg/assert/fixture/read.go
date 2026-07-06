package fixture

import "github.com/funtimecoding/soil/pkg/system"

func Read(path ...string) string {
	f := File(path...)
	defer func() {
		if e := f.Close(); e != nil {
			panic(e)
		}
	}()

	return string(system.ReadAll(f))
}
