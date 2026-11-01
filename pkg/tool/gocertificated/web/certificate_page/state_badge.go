package certificate_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"time"
)

func stateBadge(r *record.Record) gomponents.Node {
	if r.Revoked != nil {
		return badge.New(constant.RevokedState).
			Color(webConstant.BadColor).
			Node()
	}

	if r.End.Before(time.Now()) {
		return badge.New(constant.ExpiredState).
			Color(webConstant.WarnColor).
			Node()
	}

	return badge.New(constant.LiveState).Color(webConstant.OkayColor).Node()
}
