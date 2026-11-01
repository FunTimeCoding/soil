package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web/entry_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/url"
)

func (s *Server) addSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	e := entry.New()
	e.Action = r.FormValue(constant.Action)
	e.User = r.FormValue(constant.User)
	e.System = r.FormValue(constant.System)
	e.Service = r.FormValue(constant.Service)
	e.Description = r.FormValue(constant.Description)
	f := s.service.Add(e)

	if validation.Is(f) {
		s.view.RenderFragment(w, entry_page.AddForm(r.PostForm, f.Error()))

		return
	}

	if f != nil {
		s.view.Answer(w, r, s.reporter, f)

		return
	}

	s.view.RenderFragmentWithStatus(
		w,
		entry_page.AddForm(url.Values{}, ""),
		fmt.Sprintf("Entry added: %s by %s", e.Action, e.User),
		webConstant.LayoutStatusSuccess,
	)
}
