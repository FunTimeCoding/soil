package client_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maragu.dev/gomponents"
	"net/url"
)

func CreateForm(
	values url.Values,
	alert string,
) gomponents.Node {
	scopes := values.Get(constant.ScopesField)

	if scopes == "" {
		scopes = join.Space(constant.ScopeOpenIdentity, constant.ScopeOffline)
	}

	return entry.New(constant.CreatePath, "Create").
		Alert(alert).
		TextArea(
			constant.LocatorsField,
			"Redirect Locators (one per line)",
			values.Get(constant.LocatorsField),
		).
		Lines("3").
		Required().
		Text(constant.ScopesField, "Scopes", scopes).
		Node()
}
