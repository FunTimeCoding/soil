package relocation

import "github.com/dave/dst"

func NewTransplantSpec(
	declaration *dst.GenDecl,
	spec dst.Spec,
	single bool,
) *TransplantSpec {
	return &TransplantSpec{Declaration: declaration, Spec: spec, Single: single}
}
