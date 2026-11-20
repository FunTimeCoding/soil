package flag_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
	"maragu.dev/gomponents"
)

func Closed(r *overview_result.Row) gomponents.Node {
	return opening(r).Response()
}
