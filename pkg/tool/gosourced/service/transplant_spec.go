package service

import "github.com/dave/dst"

type transplantSpec struct {
	declaration *dst.GenDecl
	spec        dst.Spec
	single      bool
}
