package flag_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/subject"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func SeverityBadge(f *flag.Flag) gomponents.Node {
	if k := subject.Kind(f.Subject); k != constant.RemediationVulnerability {
		return badge.New(k).Color(webConstant.QuietColor).Node()
	}

	if f.Severity == "" {
		return badge.New("unpublished when flagged").
			Color(webConstant.QuietColor).
			Node()
	}

	color := webConstant.WarnColor

	if f.Score >= constant.CriticalScore {
		color = webConstant.BadColor
	}

	return badge.New(fmt.Sprintf("%s %.1f", f.Severity, f.Score)).
		Color(color).
		Node()
}
