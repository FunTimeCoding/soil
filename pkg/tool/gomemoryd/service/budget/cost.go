package budget

import stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"

func (b *Budget) cost(item string) int {
	if !b.empty {
		return len(stringsConstant.Unix) + len(item)
	}

	if !b.rendered {
		return b.opening + len(item)
	}

	return len(stringsConstant.Unix) + b.opening + len(item)
}
