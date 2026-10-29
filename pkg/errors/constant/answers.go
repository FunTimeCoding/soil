package constant

import (
	"github.com/funtimecoding/soil/pkg/errors/types/answer"
	"net/http"
)

const (
	ValidationClass    = "validation"
	AmbiguousClass     = "ambiguous"
	NotSelectedClass   = "not_selected"
	NotFoundClass      = "not_found"
	ConflictClass      = "conflict"
	ForbiddenClass     = "forbidden"
	UnavailableClass   = "unavailable"
	UnreachableClass   = "unreachable"
	TimeoutClass       = "timeout"
	DroppedClass       = "dropped"
	NotConfiguredClass = "not_configured"
)

var Answers = map[string]*answer.Answer{
	ValidationClass:    answer.NewUncaptured(http.StatusBadRequest),
	AmbiguousClass:     answer.NewUncaptured(http.StatusBadRequest),
	NotSelectedClass:   answer.NewUncaptured(http.StatusBadRequest),
	NotFoundClass:      answer.NewUncaptured(http.StatusNotFound),
	ConflictClass:      answer.NewUncaptured(http.StatusConflict),
	ForbiddenClass:     answer.NewUncaptured(http.StatusForbidden),
	UnavailableClass:   answer.NewUncaptured(http.StatusServiceUnavailable),
	UnreachableClass:   answer.NewCaptured(http.StatusServiceUnavailable),
	TimeoutClass:       answer.NewCaptured(http.StatusServiceUnavailable),
	DroppedClass:       answer.NewCaptured(http.StatusServiceUnavailable),
	NotConfiguredClass: answer.NewCaptured(http.StatusInternalServerError),
}
