package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/token"
	"maps"
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
			file := lookup.FindSyntaxFile(
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
		name, locals, refusal := arrivalName(
			q,
			sourcePackagePath,
			targetPackagePath,
			targetPackageName,
		)

		if refusal != "" {
			return nil, refusal
		}

		if name == nil {
			return nil, fmt.Sprintf("no available import name in %s", filename)
		}

		q.Name = name
		maps.Copy(q.Locals, locals)

		if !q.Name.Imported {
			continue
		}

		if at := shadowedSite(q, targetPackagePath); at != token.NoPos {
			return nil, lookup.TakenMessage(fileSet, q.Name.Local, at)
		}
	}

	return result, ""
}
