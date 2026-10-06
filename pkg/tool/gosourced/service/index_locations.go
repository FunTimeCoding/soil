package service

import (
	"github.com/funtimecoding/soil/pkg/source/index/xref"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result/location"
	"go/types"
	"path/filepath"
)

func indexLocations(
	directory string,
	i *xref.Index,
	declaration types.Object,
	excludeFile string,
) []*location.Location {
	target, okay := xref.Target(declaration)

	if !okay {
		return nil
	}

	var result []*location.Location

	for unit, sites := range i.Sites(target) {
		for _, s := range sites {
			full := filepath.Join(i.Directory(unit), s.File)

			if excludeFile != "" && full == excludeFile {
				continue
			}

			result = append(
				result,
				location.New(system.RelativePath(directory, full), s.Line, unit),
			)
		}
	}

	return result
}
