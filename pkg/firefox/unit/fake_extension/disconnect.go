package fake_extension

import "github.com/funtimecoding/soil/pkg/errors"

func (x *Extension) Disconnect() {
	errors.LogOnError(x.connection.CloseNow())
}
