package service

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"os"
	"path/filepath"
)

func writeMoveTarget(
	d *decoration.Set,
	plan *relocation.Plan,
	fileName string,
	transplants []dst.Decl,
) (string, error) {
	targetPath := filepath.Join(plan.MoveDirectory, fileName)

	if plan.Target != nil {
		if astFile := findSyntaxFile(
			plan.Set,
			plan.Target,
			targetPath,
		); astFile != nil {
			file, e := d.DecorateFile(plan.Set, plan.Target, astFile)

			if e != nil {
				return targetPath, e
			}

			file.Decls = append(file.Decls, transplants...)

			return targetPath, nil
		}
	}

	if _, e := os.Stat(targetPath); e == nil {
		return targetPath, fmt.Errorf(
			"target file exists but is not part of the loaded package: %s",
			targetPath,
		)
	}

	file := &dst.File{
		Name:  dst.NewIdent(plan.TargetPackageName),
		Decls: transplants,
	}

	if lines := plan.Constraints[fileName]; len(lines) > 0 {
		for _, line := range lines {
			file.Decs.Start.Append(line)
		}

		file.Decs.Start.Append("\n")
	}

	d.Files[targetPath] = file
	d.PackagePaths[file] = plan.TargetPackagePath

	return targetPath, nil
}
