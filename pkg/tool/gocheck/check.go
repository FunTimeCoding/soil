package gocheck

import (
	"github.com/funtimecoding/soil/pkg/source/build_tag"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
)

func check(
	o *option.Check,
	local bool,
) (*check_result.Result, error) {
	if !local {
		r, delegated, e := Delegate(o)

		if e != nil || delegated {
			return r, e
		}
	}

	build_tag.Memoize()

	return Run(o)
}
