package flag_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
	"maragu.dev/gomponents"
)

func Closed(
	r *overview_result.Row,
	c *affected_count.Count,
) gomponents.Node {
	return opening(r, c).Response()
}
