package unguarded_peer

import "go/types"

type Reference struct {
	Guarded   bool
	Enclosing types.Object
}
