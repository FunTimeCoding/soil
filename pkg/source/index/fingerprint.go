package index

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/types"
	"sort"
)

func Fingerprint(p *types.Package) string {
	var lines []string
	scope := p.Scope()

	for _, name := range scope.Names() {
		o := scope.Lookup(name)

		if !o.Exported() {
			continue
		}

		lines = append(lines, types.ObjectString(o, qualifier))
		typeName, isTypeName := o.(*types.TypeName)

		if !isTypeName {
			continue
		}

		set := types.NewMethodSet(types.NewPointer(typeName.Type()))

		if _, isFace := typeName.Type().Underlying().(*types.Interface); isFace {
			set = types.NewMethodSet(typeName.Type())
		}

		for i := range set.Len() {
			if f, okay := set.At(i).Obj().(*types.Func); okay && f.Exported() {
				lines = append(lines, join.Space(name, f.Name(), Signature(f)))
			}
		}
	}

	sort.Strings(lines)
	h := sha256.New()

	for _, l := range lines {
		h.Write([]byte(l))
		h.Write([]byte{0})
	}

	return hex.EncodeToString(h.Sum(nil))
}
