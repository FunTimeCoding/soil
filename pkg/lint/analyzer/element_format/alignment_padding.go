package element_format

import "github.com/dave/dst"

func AlignmentPadding(
	parent *dst.CompositeLit,
	keyLength int,
) int {
	longest := 0

	for _, el := range parent.Elts {
		kv, okay := el.(*dst.KeyValueExpr)

		if !okay {
			continue
		}

		l := KeyWidth(kv.Key)

		if l > longest {
			longest = l
		}
	}

	if longest <= keyLength {
		return 0
	}

	return longest - keyLength
}
