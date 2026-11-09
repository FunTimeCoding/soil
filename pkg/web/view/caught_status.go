package view

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"net/http"
)

func caughtStatus(caught any) int {
	if e, okay := caught.(error); okay {
		if a := classify.Answer(e); a != nil {
			return a.Status
		}
	}

	return http.StatusInternalServerError
}
