package coverage_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"maragu.dev/gomponents"
)

func Closed(e *coverage.Server) gomponents.Node {
	return opening(e).Response()
}
