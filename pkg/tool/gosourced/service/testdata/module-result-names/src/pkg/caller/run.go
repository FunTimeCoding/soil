package caller

import "example/lib/v2"

func Run() error {
	if e := lib.Encode(1); e != nil {
		return e
	}

	return lib.Close()
}
