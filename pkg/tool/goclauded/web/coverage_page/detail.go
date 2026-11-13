package coverage_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"maragu.dev/gomponents"
)

func Detail(e *coverage.Server) gomponents.Node {
	return opening(e).Detail(len(cells(e)), tools(e)).Response()
}
