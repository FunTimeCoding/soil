package scan

import (
	"github.com/funtimecoding/soil/pkg/parse"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/matrix"
	"path/filepath"
)

func scanFrontend(
	v *virtual_file_system.System,
	s *Service,
) *matrix.Frontend {
	web := filepath.Join(s.Path, constant.WebDirectory)
	construct := parseWebFile(v, filepath.Join(web, constant.NewFileName))

	if construct == nil ||
		!parse.HasCall(construct, constant.LayoutImport, "New") {
		return nil
	}

	f := matrix.NewFrontend(
		s.Name,
		s.Repo,
		constant.UnknownTheme,
		v.Has(filepath.Join(web, constant.FaviconFile)),
	)
	name, _ := parse.ImportName(construct, constant.LayoutImport)
	layouts := parse.FindCalls(construct, name, "New")

	if len(layouts) > 0 && len(layouts[0].Args) > 1 {
		f.Theme = themeName(layouts[0].Args[1])
	}

	f.Style = len(parse.FindMethods(construct, constant.StyleMethod)) > 0
	f.Palette = len(parse.FindMethods(construct, constant.ItemsMethod))+
		len(parse.FindMethods(construct, constant.CommandsMethod)) > 0
	f.Live = len(parse.FindMethods(construct, constant.LiveMethod)) > 0
	items := parse.FindMethods(construct, constant.ItemsMethod)

	if len(items) > 0 {
		f.Items = len(items[0].Args)
	}

	mount := parseWebFile(v, filepath.Join(web, constant.MountFileName))

	if mount == nil {
		return f
	}

	if mountsView(mount) {
		f.PaletteRoute = true
		f.FaviconRoute = true
	}

	return f
}
