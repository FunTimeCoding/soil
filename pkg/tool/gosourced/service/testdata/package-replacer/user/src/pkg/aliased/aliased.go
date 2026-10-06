package aliased

import figure "other.test/lib/shape"

func Make() *figure.Shape {
	return figure.New()
}
