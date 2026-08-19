package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func createAuthorityForm(r *http.Request) gomponents.Node {
	return html.Form(
		html.Method(http.MethodPost),
		html.Action(constant.CreateAuthorityPath),
		field(r, "name", "Name", "root, cluster or host"),
		field(r, constant.KindParameter, "Kind", "root or intermediate"),
		field(r, constant.CommonNameParameter, "Common name", ""),
		field(r, constant.CountryParameter, "Country", "root only"),
		field(r, constant.ProvinceParameter, "Province", "root only"),
		field(r, constant.OrganizationParameter, "Organization", "root only"),
		field(
			r,
			constant.DomainParameter,
			"Permitted domains",
			"comma separated, intermediates only",
		),
		field(
			r,
			constant.AddressParameter,
			"Permitted addresses",
			"comma separated CIDR, intermediates only",
		),
		html.Button(html.Type("submit"), gomponents.Text("Create")),
	)
}
