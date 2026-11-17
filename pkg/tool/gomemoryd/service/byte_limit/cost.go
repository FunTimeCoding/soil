package byte_limit

import stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"

func (l *Limit) cost(item string) int {
	if !l.empty {
		return len(stringsConstant.Unix) + len(item)
	}

	if !l.rendered {
		return l.opening + len(item)
	}

	return len(stringsConstant.Unix) + l.opening + len(item)
}
