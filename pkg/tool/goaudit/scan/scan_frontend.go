package scan

import (
	"github.com/funtimecoding/soil/pkg/parse"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
)

func scanFrontend(
	v *virtual_file_system.System,
	s *Service,
) *Frontend {
	web := filepath.Join(s.Path, constant.WebDirectory)
	construct := parseWebFile(v, filepath.Join(web, constant.NewFileName))

	if construct == nil ||
		!parse.HasCall(construct, constant.LayoutImport, "New") {
		return nil
	}

	f := &Frontend{
		Name:    s.Name,
		Repo:    s.Repo,
		Theme:   constant.UnknownTheme,
		Favicon: v.Has(filepath.Join(web, constant.FaviconFile)),
	}
	theme := parse.FindMethods(construct, constant.ThemeMethod)

	if len(theme) > 0 {
		f.Theme = themeName(theme[0])
	}

	f.Style = len(parse.FindMethods(construct, constant.StyleMethod)) > 0
	f.Palette = len(parse.FindMethods(construct, constant.PaletteMethod)) > 0
	f.Live = len(parse.FindMethods(construct, constant.LiveMethod)) > 0
	items := parse.FindMethods(construct, constant.ItemsMethod)

	if len(items) > 0 {
		f.Items = len(items[0].Args)
	}

	mount := parseWebFile(v, filepath.Join(web, constant.MountFileName))

	if mount == nil {
		return f
	}

	for _, c := range parse.FindMethods(mount, constant.RouteMethod) {
		if len(c.Args) == 0 {
			continue
		}

		route, okay := parse.StringValue(c.Args[0])

		if !okay {
			continue
		}

		if route == constant.PaletteRoute {
			f.PaletteRoute = true
		}

		if route == constant.FaviconRoute {
			f.FaviconRoute = true
		}
	}

	return f
}
