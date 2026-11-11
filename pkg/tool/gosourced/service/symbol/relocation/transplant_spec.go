package relocation

import "github.com/dave/dst"

type TransplantSpec struct {
	Declaration *dst.GenDecl
	Spec        dst.Spec
	Single      bool
}
