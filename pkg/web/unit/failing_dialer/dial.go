package failing_dialer

import (
	"context"
	"net"
)

func (d *Dialer) dial(
	context.Context,
	string,
	string,
) (net.Conn, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.attempts++

	return nil, d.failure
}
