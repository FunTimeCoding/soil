package kind

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

type Kind struct {
	Name     string
	Value    func() any
	Extract  func(*packages.Package) any
	External func(*types.Package) any
}
