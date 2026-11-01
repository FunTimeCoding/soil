package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/types/change"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web/entry_page"
	"net/http"
	"time"
)

func (s *Server) editSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	e, f := s.entryFromQuery(r)

	if f != nil {
		s.view.Answer(w, r, s.reporter, f)

		return
	}

	errors.PanicOnError(r.ParseForm())
	action := r.FormValue(constant.Action)
	user := r.FormValue(constant.User)
	system := r.FormValue(constant.System)
	service := r.FormValue(constant.Service)
	description := r.FormValue(constant.Description)
	c := change.New()
	c.Action = &action
	c.User = &user
	c.System = &system
	c.Service = &service
	c.Description = &description

	if v := r.FormValue(constant.Timestamp); v != "" {
		t, g := time.Parse(constant.FilterFormat, v)

		if g != nil {
			c.Apply(e)
			s.view.RenderFragment(
				w,
				entry_page.EditRow(e, fmt.Sprintf("invalid timestamp: %s", v)),
			)

			return
		}

		c.Timestamp = &t
	}

	result, g := s.service.Update(e.Identifier, c)

	if validation.Is(g) {
		c.Apply(e)
		s.view.RenderFragment(w, entry_page.EditRow(e, g.Error()))

		return
	}

	if g != nil {
		s.view.Answer(w, r, s.reporter, g)

		return
	}

	s.view.RenderFragment(w, entry_page.Detail(result))
}
