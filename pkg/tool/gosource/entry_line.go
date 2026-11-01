package gosource

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosource/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
)

func entryLine(e client.Entry) string {
	switch {
	case e.From != nil && e.To != nil:
		return join.Space(
			e.Class,
			locationText(e.To),
			constant.SourceArrow,
			locationText(e.From),
		)
	case e.To != nil:
		return join.Space(e.Class, locationText(e.To))
	default:
		return join.Space(e.Class, locationText(e.From))
	}
}
