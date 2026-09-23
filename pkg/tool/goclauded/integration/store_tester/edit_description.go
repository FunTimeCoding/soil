package store_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument/edit_session"

func (o *Tester) EditDescription(
	identifier string,
	description string,
) {
	o.Edit(identifier, edit_session.New().WithDescription(description))
}
