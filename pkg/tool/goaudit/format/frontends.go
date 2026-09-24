package format

import (
	"github.com/funtimecoding/soil/pkg/console/table"
	"github.com/funtimecoding/soil/pkg/integers"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
)

func Frontends(frontends []*scan.Frontend) string {
	t := table.New(
		"FRONTEND",
		"REPO",
		"THEME",
		"STYLE",
		"PALETTE",
		"ITEMS",
		"LIVE",
		"FAVICON",
	)

	for _, f := range frontends {
		t.Add(
			f.Name,
			f.Repo,
			f.Theme,
			mark(f.Style),
			pairMark(f.Palette, f.PaletteRoute),
			integers.ToString(f.Items),
			mark(f.Live),
			pairMark(f.Favicon, f.FaviconRoute),
		)
	}

	return t.Render()
}
