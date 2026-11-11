package caller

import "example/lib/v2"

func Run(in []byte) error {
	var out map[string]string

	return lib.Decode(in, &out)
}
