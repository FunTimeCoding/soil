package caller

import "example/pkg/target"

func Run() target.Mode {
	if target.Off == 0 {
		return target.On
	}

	return target.Off
}
