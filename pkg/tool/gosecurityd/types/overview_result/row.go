package overview_result

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/status_result"
)

type Row struct {
	Flag   *flag.Flag
	Debian *status_result.Result
	Unread error
	Fix    string
}
