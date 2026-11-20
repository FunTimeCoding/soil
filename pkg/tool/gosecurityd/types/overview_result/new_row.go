package overview_result

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/status_result"
)

func NewRow(
	f *flag.Flag,
	d *status_result.Result,
	unread error,
	fix string,
) *Row {
	return &Row{Flag: f, Debian: d, Unread: unread, Fix: fix}
}
