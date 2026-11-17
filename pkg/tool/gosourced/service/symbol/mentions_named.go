package symbol

import "go/types"

func mentionsNamed(
	t types.Type,
	wanted map[*types.Named]bool,
) bool {
	switch u := t.(type) {
	case *types.Named:
		return wanted[u.Origin()]
	case *types.Pointer:
		return mentionsNamed(u.Elem(), wanted)
	case *types.Slice:
		return mentionsNamed(u.Elem(), wanted)
	case *types.Array:
		return mentionsNamed(u.Elem(), wanted)
	case *types.Chan:
		return mentionsNamed(u.Elem(), wanted)
	case *types.Map:
		return mentionsNamed(u.Key(), wanted) || mentionsNamed(u.Elem(), wanted)
	}

	return false
}
