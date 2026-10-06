package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"go/token"
)

func planQualifications(
	fileSet *token.FileSet,
	references []relocation.QualifiedReference,
	sourcePackagePath string,
	targetPackagePath string,
	targetPackageName string,
) (map[string]*relocation.FileQualification, string) {
	result := make(map[string]*relocation.FileQualification)

	for _, f := range references {
		position := fileSet.Position(f.Reference.Ident.Pos())
		q, exists := result[position.Filename]

		if !exists {
			file := findSyntaxFile(
				fileSet,
				f.Reference.Package,
				position.Filename,
			)

			if file == nil {
				continue
			}

			q = relocation.NewFileQualification(
				file,
				f.Reference.Package,
				sourcePackagePath,
			)
			result[position.Filename] = q
		}

		q.Idents[f.Reference.Ident] = f.NewName
		q.Positions = append(
			q.Positions,
			relocation.QualifiedPosition{
				Position: position,
				OldName:  f.Reference.Ident.Name,
				NewName:  f.NewName,
			},
		)
	}

	for filename, q := range result {
		q.Name = chooseImportName(q.File, targetPackagePath, targetPackageName)

		if q.Name == nil {
			return nil, filename
		}
	}

	return result, ""
}
