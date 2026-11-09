package symbol

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func methodFamily(
	all []*packages.Package,
	m *types.Func,
) ([]*types.Func, []string) {
	name := m.Name()
	owner := m.Type().(*types.Signature).Recv().Type().Underlying()
	active := []*types.Interface{owner.(*types.Interface)}
	members := []*types.Func{m}
	joined := map[*types.Func]bool{m: true}
	joinedType := map[*types.Named]bool{}
	var concrete []*types.Named
	var refusals []string

	for changed := true; changed; {
		changed = false

		for _, t := range namedTypes(all) {
			if joinedType[t] {
				continue
			}

			i, isInterface := t.Underlying().(*types.Interface)

			if isInterface {
				d := explicitMethod(i, name)

				if d == nil || joined[d] || !anySatisfies(concrete, i) {
					continue
				}

				joinedType[t] = true
				joined[d] = true
				members = append(members, d)
				active = append(active, i)
				changed = true

				continue
			}

			o, _, _ := types.LookupFieldOrMethod(
				types.NewPointer(t),
				true,
				m.Pkg(),
				name,
			)
			f, isFunction := o.(*types.Func)

			if !isFunction {
				continue
			}

			if t.TypeParams().Len() > 0 {
				joinedType[t] = true
				refusals = append(
					refusals,
					fmt.Sprintf(
						"%s is generic and has %s - its satisfaction can't be checked",
						types.TypeString(t, nil),
						name,
					),
				)

				continue
			}

			if !implementsAny(t, active) {
				continue
			}

			joinedType[t] = true
			concrete = append(concrete, t)

			if !joined[f] {
				joined[f] = true
				members = append(members, f)
			}

			changed = true
		}
	}

	return members, refusals
}
