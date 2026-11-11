package caller

import "example/lib/v2"

func Run() string {
	return lib.Next().Value
}
