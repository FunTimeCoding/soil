package service

import "go/token"

func planQualifications(
	fileSet *token.FileSet,
	references []qualifiedReference,
	sourcePackagePath string,
	targetPackagePath string,
	targetPackageName string,
) (map[string]*fileQualification, string) {
	result := make(map[string]*fileQualification)

	for _, f := range references {
		position := fileSet.Position(f.reference.Ident.Pos())
		q, exists := result[position.Filename]

		if !exists {
			file := findSyntaxFile(
				fileSet,
				f.reference.Package,
				position.Filename,
			)

			if file == nil {
				continue
			}

			q = newFileQualification(
				file,
				f.reference.Package,
				sourcePackagePath,
			)
			result[position.Filename] = q
		}

		q.idents[f.reference.Ident] = f.newName
		q.positions = append(
			q.positions,
			qualifiedPosition{
				position: position,
				oldName:  f.reference.Ident.Name,
				newName:  f.newName,
			},
		)
	}

	for filename, q := range result {
		q.name = chooseImportName(
			q.file,
			targetPackagePath,
			targetPackageName,
		)

		if q.name == nil {
			return nil, filename
		}
	}

	return result, ""
}
