package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/web/form"
	"net/http"
)

func (s *Server) createAuthoritySubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	b := &server.AuthorityBody{
		Name: r.PostFormValue("name"),
		Kind: server.AuthorityKind(
			r.PostFormValue(constant.KindParameter),
		),
		CommonName: r.PostFormValue(constant.CommonNameParameter),
	}
	optionalText(&b.Country, r.PostFormValue(constant.CountryParameter))
	optionalText(&b.Province, r.PostFormValue(constant.ProvinceParameter))
	optionalText(
		&b.Organization,
		r.PostFormValue(constant.OrganizationParameter),
	)
	optionalList(&b.PermittedDomain, r.PostFormValue(constant.DomainParameter))
	optionalList(
		&b.PermittedAddress,
		r.PostFormValue(constant.AddressParameter),
	)

	if _, e := s.service.CreateAuthority(b); e != nil {
		form.Redirect(w, r, constant.CreateAuthorityPath, e.Error(), r.PostForm)

		return
	}

	http.Redirect(w, r, constant.AuthoritiesPath, http.StatusSeeOther)
}
