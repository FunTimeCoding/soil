package caller

import "example/pkg/target"

func Main(values []int) string {
	v := "v1"
	o := &target.Option{Name: "x", Version: v}
	label := o.Version

	return target.Run(o, label) +
		target.Log("a", 1, 2) +
		target.Log("b", values...)
}
