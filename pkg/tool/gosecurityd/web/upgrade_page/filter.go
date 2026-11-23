package upgrade_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/system_page"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"maps"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Filter(
	s *system_selection.Selection,
	release string,
	kernel bool,
	releases []string,
) gomponents.Node {
	link := func(r string, k bool) string {
		q := s.Query(0)
		maps.Copy(q, Kept(r, k))

		return system_page.Query(constant.UpgradesPath, q)
	}
	options := append([]string{""}, releases...)
	byRelease := make([]gomponents.Node, 0, len(options))

	for _, r := range options {
		label := r

		if r == "" {
			label = "All releases"
		}

		c := chip.New(label).Link(link(r, kernel))

		if r == release {
			c.Current()
		}

		byRelease = append(byRelease, c.Node())
	}

	byKind := make([]gomponents.Node, 0, 2)

	for _, k := range []bool{false, true} {
		label := "Packages"

		if k {
			label = "Kernel"
		}

		c := chip.New(label).Link(link(release, k))

		if k == kernel {
			c.Current()
		}

		byKind = append(byKind, c.Node())
	}

	return gomponents.Group(
		[]gomponents.Node{html.Nav(byRelease...), html.Nav(byKind...)},
	)
}
