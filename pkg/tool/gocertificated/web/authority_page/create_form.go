package authority_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/web/component/choice"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maragu.dev/gomponents"
	"net/url"
)

func CreateForm(
	values url.Values,
	alert string,
) gomponents.Node {
	return entry.New(constant.CreateAuthorityPath, "Create").
		Alert(alert).
		Text("name", "Name", values.Get("name")).
		Placeholder("root, cluster or host").
		Select(
			constant.KindParameter,
			"Kind",
			choice.FromValues(
				[]string{
					string(server.AuthorityKindRoot),
					string(server.AuthorityKindIntermediate),
				},
			),
			values.Get(constant.KindParameter),
		).
		Text(
			constant.CommonNameParameter,
			"Common name",
			values.Get(constant.CommonNameParameter),
		).
		Text(
			constant.CountryParameter,
			"Country",
			values.Get(constant.CountryParameter),
		).
		Placeholder("root only").
		Text(
			constant.ProvinceParameter,
			"Province",
			values.Get(constant.ProvinceParameter),
		).
		Placeholder("root only").
		Text(
			constant.OrganizationParameter,
			"Organization",
			values.Get(constant.OrganizationParameter),
		).
		Placeholder("root only").
		Text(
			constant.DomainParameter,
			"Permitted domains",
			values.Get(constant.DomainParameter),
		).
		Placeholder("comma separated, intermediates only").
		Text(
			constant.AddressParameter,
			"Permitted addresses",
			values.Get(constant.AddressParameter),
		).
		Placeholder("comma separated CIDR, intermediates only").
		Node()
}
