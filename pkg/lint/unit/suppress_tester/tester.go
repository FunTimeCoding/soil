package suppress_tester

import (
	"go/ast"
	"go/token"
	"testing"
)

type Tester struct {
	t       *testing.T
	fileSet *token.FileSet
	file    *ast.File
}
