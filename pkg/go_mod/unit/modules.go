package unit

import (
	"github.com/funtimecoding/soil/pkg/go_mod/module"
	"github.com/funtimecoding/soil/pkg/go_mod/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/time/day"
	"time"
)

func modules(stream string) []*module.Module {
	var result []*module.Module

	for _, r := range notation.MustDecodeStream[response.Module](stream) {
		m := module.New(r, day.New(2026, time.October, 11), 12)
		m.Validate()
		result = append(result, m)
	}

	return result
}
