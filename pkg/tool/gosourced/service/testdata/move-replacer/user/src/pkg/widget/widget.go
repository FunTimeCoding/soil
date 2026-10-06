package widget

import "other.test/lib/source"

type Holder struct {
	source.Widget
}

func Make() int {
	w := &source.Widget{Size: 1}

	return w.Run() + source.Limit
}
