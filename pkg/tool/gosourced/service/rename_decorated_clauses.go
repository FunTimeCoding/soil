package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"go/ast"
)

func renameDecoratedClauses(
	decorations *decoration.Set,
	modified map[string]*ast.File,
	oldName string,
	newName string,
) {
	testOld := join.Empty(oldName, "_test")
	testNew := join.Empty(newName, "_test")

	for filename := range modified {
		file := decorations.Files[filename]

		if file == nil {
			continue
		}

		if file.Name.Name == oldName {
			file.Name.Name = newName
		}

		if file.Name.Name == testOld {
			file.Name.Name = testNew
		}
	}
}
