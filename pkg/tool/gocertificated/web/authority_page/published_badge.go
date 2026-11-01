package authority_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func publishedBadge(r *record.Record) gomponents.Node {
	if r.Published == nil {
		return badge.New("pending").Color(webConstant.WarnColor).Node()
	}

	return badge.New("yes").Color(webConstant.OkayColor).Node()
}
