package directory

import (
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/go-ldap/ldap/v3"
	"strings"
)

func (c *Client) searchRequest(
	filter string,
	value string,
) *ldap.SearchRequest {
	return ldap.NewSearchRequest(
		c.base,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		strings.ReplaceAll(filter, "%s", ldap.EscapeFilter(value)),
		[]string{
			constant.UniqueAttribute,
			constant.AccountAttribute,
			constant.MailAttribute,
			constant.NameAttribute,
		},
		nil,
	)
}
