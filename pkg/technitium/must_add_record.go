package technitium

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/technitium/record"
)

func (c *Client) MustAddRecord(
	domain string,
	recordType string,
	value string,
) *record.Record {
	result, e := c.AddRecord(domain, recordType, value)
	errors.PanicOnError(e)

	return result
}
