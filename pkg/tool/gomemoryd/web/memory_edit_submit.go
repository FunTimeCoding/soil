package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/memory_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) memoryEditSubmit(
	w http.ResponseWriter,
	q *http.Request,
) {
	v, okay := s.visibleMemory(w, q)

	if !okay {
		return
	}

	errors.PanicOnError(q.ParseForm())
	o := save_option.New()
	o.Name = q.FormValue(constant.MemoryName)
	o.Description = q.FormValue(constant.Description)
	o.Content = q.FormValue(constant.Content)
	o.Source = constant.WebSource
	_, rewritten, e := s.service.EditMemory(v.Memory.Identifier, o)

	if validation.Is(e) || conflict.Is(e) || forbidden.Is(e) {
		v.Memory.Name = o.Name
		v.Memory.Description = o.Description
		v.Memory.Content = o.Content
		s.view.RenderFragment(w, memory_page.New(v).Editor(e.Error()))

		return
	}

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	saved, f := s.service.MemoryView(v.Memory.Identifier)

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return
	}

	status := "Saved"

	if len(rewritten) > 0 {
		status = fmt.Sprintf(
			"Saved; citations rewritten in %d memories",
			len(rewritten),
		)
	}

	s.view.RenderStatusFragment(
		w,
		memory_page.New(saved).Opened(),
		status,
		webConstant.LayoutStatusSuccess,
	)
}
