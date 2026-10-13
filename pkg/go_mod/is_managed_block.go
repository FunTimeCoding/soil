package go_mod

import (
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"golang.org/x/mod/modfile"
)

func isManagedBlock(b *modfile.LineBlock) bool {
	if len(b.Token) == 0 || b.Token[0] != constant.ReplaceKeyword {
		return false
	}

	for _, c := range b.Before {
		if c.Token == constant.ManagedMarker {
			return true
		}
	}

	return false
}
