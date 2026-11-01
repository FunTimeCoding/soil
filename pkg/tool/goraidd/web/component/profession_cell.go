package component

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func ProfessionCell(profession string) gomponents.Node {
	return html.Td(
		html.Img(
			html.Src(
				fmt.Sprintf(
					"%sicons/%s.png",
					constant.StaticPath,
					strings.ToLower(profession),
				),
			),
			html.Width("20"),
			html.Height("20"),
		),
		gomponents.Textf(" %s", profession),
	)
}
