package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/web/group_page"
	"github.com/funtimecoding/soil/pkg/web/component"
	"github.com/funtimecoding/soil/pkg/web/form"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) group(
	w http.ResponseWriter,
	r *http.Request,
) {
	found, e := s.service.Groups()
	errors.PanicOnError(e)
	values := r.URL.Query()
	s.view.RenderPage(
		w,
		constant.GroupTitle,
		constant.GroupPath,
		html.H1(gomponents.Text(constant.GroupTitle)),
		component.Alert(form.PageError(r)),
		group_page.Table(found),
		group_page.CreateForm(values, form.Failed(r, constant.CreateGroupForm)),
		group_page.MemberForm(values, form.Failed(r, constant.MemberForm)),
	)
}
