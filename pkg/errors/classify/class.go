package classify

import (
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func Class(e error) string {
	switch {
	case unavailable.Is(e):
		return constant.UnavailableClass
	case not_configured.Is(e):
		return constant.NotConfiguredClass
	case not_found.Is(e):
		return constant.NotFoundClass
	case ambiguous.Is(e):
		return constant.AmbiguousClass
	case not_selected.Is(e):
		return constant.NotSelectedClass
	case conflict.Is(e):
		return constant.ConflictClass
	case validation.Is(e):
		return constant.ValidationClass
	case timeout.Is(e):
		return constant.TimeoutClass
	case dropped.Is(e):
		return constant.DroppedClass
	case unreachable.Is(e):
		return constant.UnreachableClass
	default:
		return ""
	}
}
