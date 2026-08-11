package format

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"strings"
)

func Frontends(frontends []*scan.Frontend) string {
	var b strings.Builder
	t := newTable(
		[]string{
			"FRONTEND",
			"REPO",
			"THEME",
			"STYLE",
			"PALETTE",
			"ITEMS",
			"LIVE",
			"FAVICON",
		},
	)

	for _, f := range frontends {
		t.addRow(
			[]string{
				f.Name,
				f.Repo,
				f.Theme,
				mark(f.Style),
				pairMark(f.Palette, f.PaletteRoute),
				fmt.Sprintf("%d", f.Items),
				mark(f.Live),
				pairMark(f.Favicon, f.FaviconRoute),
			},
		)
	}

	b.WriteString(t.render())

	return b.String()
}
