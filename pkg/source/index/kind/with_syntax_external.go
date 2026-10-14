package kind

import "golang.org/x/tools/go/packages"

func (k *Kind) WithSyntaxExternal(f func(*packages.Package) any) *Kind {
	k.ExternalSyntax = f

	return k
}
