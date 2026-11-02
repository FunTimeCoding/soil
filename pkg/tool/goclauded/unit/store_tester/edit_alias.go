package store_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/argument/edit_session"

func (o *Tester) EditAlias(
	identifier string,
	alias string,
) {
	o.Edit(identifier, edit_session.New().WithAlias(alias))
}
