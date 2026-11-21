package upgrade_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_group"
	"net/url"
)

func fragment(
	path string,
	g *upgrade_group.Group,
) string {
	return join.Empty(
		path,
		"?",
		url.Values{
			constant.ReleaseParameter: {g.Release},
			constant.SourceParameter:  {g.Source},
		}.Encode(),
	)
}
