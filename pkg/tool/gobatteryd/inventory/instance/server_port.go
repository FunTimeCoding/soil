package instance

import nutConstant "github.com/funtimecoding/soil/pkg/nut/constant"

func (i *Instance) ServerPort() int {
	if i.Port == 0 {
		return nutConstant.Port
	}

	return i.Port
}
