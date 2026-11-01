package component

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func LabelChips(labels []label.Label) gomponents.Node {
	if len(labels) == 0 {
		return gomponents.Group(nil)
	}

	chips := make([]gomponents.Node, 0, len(labels))

	for _, l := range labels {
		chips = append(
			chips,
			chip.New(fmt.Sprintf("%s: %s", l.Key, l.Value)).Node(),
		)
	}

	return html.Div(
		html.Class(webConstant.ClusterClass),
		gomponents.Group(chips),
	)
}
