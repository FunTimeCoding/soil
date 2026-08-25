package service

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"os"
	"path/filepath"
)

func writeMoveTarget(
	d *decoration.Set,
	plan *movePlan,
	fileName string,
	transplants []dst.Decl,
) (string, error) {
	targetPath := filepath.Join(plan.moveDirectory, fileName)

	if plan.target != nil {
		if astFile := findSyntaxFile(
			plan.set,
			plan.target,
			targetPath,
		); astFile != nil {
			file, e := d.DecorateFile(plan.set, plan.target, astFile)

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
		Name:  dst.NewIdent(plan.targetPackageName),
		Decls: transplants,
	}

	if lines := plan.constraints[fileName]; len(lines) > 0 {
		for _, line := range lines {
			file.Decs.Start.Append(line)
		}

		file.Decs.Start.Append("\n")
	}

	d.Files[targetPath] = file
	d.PackagePaths[file] = plan.targetPackagePath

	return targetPath, nil
}
