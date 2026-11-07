package policy

import (
	"mvdan.cc/sh/v3/syntax"
	"path/filepath"
)

func (w *Walk) changeDirectory(arguments []*syntax.Word) bool {
	previous := w.base
	w.base = ""
	list, options := operands(arguments)

	if len(list) == 1 && len(options) == 0 {
		if target, okay := w.resolve(list[0]); okay && target != "" {
			if !filepath.IsAbs(target) && previous != "" {
				target = filepath.Join(previous, target)
			}

			if settled, settledOkay := settle(target); settledOkay {
				w.base = settled
			}
		}
	}

	if w.base == "" || w.base != w.origin {
		w.moved = true
	}

	return true
}
