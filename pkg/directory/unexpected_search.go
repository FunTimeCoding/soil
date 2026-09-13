package directory

import (
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/go-ldap/ldap/v3"
)

func unexpectedSearch(e error) error {
	if ldap.IsErrorWithCode(e, ldap.LDAPResultNoSuchObject) {
		return not_found.Format("%s base missing", constant.Subject)
	}

	return e
}
