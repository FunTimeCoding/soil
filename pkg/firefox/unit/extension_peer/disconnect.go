package extension_peer

import "github.com/funtimecoding/soil/pkg/errors"

func (x *Extension) Disconnect() {
	errors.LogOnError(x.connection.CloseNow())
}
