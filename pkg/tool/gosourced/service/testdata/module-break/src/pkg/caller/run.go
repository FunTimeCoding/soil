package caller

import "example/lib/v2"

func Run() string {
	w := lib.Make("alfa")
	w.Label = lib.Gone()

	return w.Name()
}
