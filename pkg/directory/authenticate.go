package directory

import (
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/directory/types/entry"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (c *Client) Authenticate(
	account string,
	password string,
) (*entry.Entry, error) {
	connection, e := c.connect()

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(connection)
	result, e := connection.Search(c.searchRequest(c.userFilter, account))

	if e != nil {
		return nil, unexpectedSearch(e)
	}

	if len(result.Entries) == 0 {
		return nil, not_found.New(constant.Subject, account)
	}

	if len(result.Entries) > 1 {
		return nil, ambiguous.Format(
			"%s matched %d entries: %s",
			constant.Subject,
			len(result.Entries),
			account,
		)
	}

	found := result.Entries[0]

	if e = connection.Bind(found.DN, password); e != nil {
		return nil, validation.New("invalid credentials")
	}

	return entry.New(
		found.GetAttributeValue(constant.UniqueAttribute),
		found.GetAttributeValue(constant.AccountAttribute),
		found.GetAttributeValue(constant.MailAttribute),
		found.GetAttributeValue(constant.NameAttribute),
		found.DN,
	), nil
}
