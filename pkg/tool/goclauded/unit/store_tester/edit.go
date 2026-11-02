package store_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/argument/edit_session"

func (o *Tester) Edit(
	identifier string,
	a *edit_session.Session,
) {
	o.EditSession(identifier, a)
}
