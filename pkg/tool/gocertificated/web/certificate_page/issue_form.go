package certificate_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/web/component/choice"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maragu.dev/gomponents"
	"net/url"
)

func IssueForm(
	authorities []record.Record,
	values url.Values,
	alert string,
) gomponents.Node {
	var intermediates []string

	for _, a := range authorities {
		if a.Kind == string(server.AuthorityKindIntermediate) {
			intermediates = append(intermediates, a.Name)
		}
	}

	return entry.New(constant.IssueCertificatePath, "Issue").
		Alert(alert).
		Select(
			constant.AuthorityParameter,
			"Authority",
			choice.FromValues(intermediates),
			values.Get(constant.AuthorityParameter),
		).
		Select(
			constant.KindParameter,
			"Kind",
			choice.FromValues(
				[]string{
					string(server.LeafKindServer),
					string(server.LeafKindClient),
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
			constant.HostParameter,
			"Hosts",
			values.Get(constant.HostParameter),
		).
		Placeholder("comma separated names and addresses").
		Node()
}
