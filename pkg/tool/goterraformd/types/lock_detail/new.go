package lock_detail

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
)

func New(l *lease.Lease) *Detail {
	if l == nil {
		return nil
	}

	v := l.Raw.Annotations[constant.LockAnnotation]

	if v == "" {
		return nil
	}

	result := &Detail{}

	if notation.Decode(v, result) != nil {
		return nil
	}

	return result
}
