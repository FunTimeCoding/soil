package caller

import "example/lib/v2"

func Run(k lib.Kind) lib.Kind {
	n := lib.Node{Kind: k}

	return n.Kind
}
