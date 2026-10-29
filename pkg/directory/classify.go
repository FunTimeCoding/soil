package directory

import (
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/go-ldap/ldap/v3"
)

func classify(
	e error,
	identifier string,
) error {
	if e == nil {
		return nil
	}

	switch {
	case ldap.IsErrorWithCode(e, ldap.LDAPResultEntryAlreadyExists):
		return conflict.Exists(constant.Subject, identifier)
	case ldap.IsErrorWithCode(e, ldap.LDAPResultNoSuchObject):
		return not_found.New(constant.Subject, identifier)
	case ldap.IsErrorWithCode(e, ldap.LDAPResultInvalidCredentials),
		ldap.IsErrorWithCode(e, ldap.LDAPResultInsufficientAccessRights):

		return forbidden.Format("%s %s: %s", constant.Subject, identifier, e)
	case ldap.IsErrorWithCode(e, ldap.LDAPResultConstraintViolation),
		ldap.IsErrorWithCode(e, ldap.LDAPResultInvalidDNSyntax),
		ldap.IsErrorWithCode(e, ldap.LDAPResultObjectClassViolation):

		return validation.New("%s %s: %s", constant.Subject, identifier, e)
	case ldap.IsErrorWithCode(e, ldap.LDAPResultBusy),
		ldap.IsErrorWithCode(e, ldap.LDAPResultUnavailable),
		ldap.IsErrorWithCode(e, ldap.LDAPResultUnwillingToPerform):

		return unavailable.Format("%s %s: %s", constant.Subject, identifier, e)
	default:
		return e
	}
}
