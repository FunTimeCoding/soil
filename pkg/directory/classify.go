package directory

import (
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/go-ldap/ldap/v3"
)

func classify(
	e error,
	identifier string,
) error {
	if e == nil {
		return nil
	}

	if ldap.IsErrorWithCode(e, ldap.LDAPResultEntryAlreadyExists) {
		return conflict.Exists(constant.Subject, identifier)
	}

	if ldap.IsErrorWithCode(e, ldap.LDAPResultNoSuchObject) {
		return not_found.New(constant.Subject, identifier)
	}

	return e
}
