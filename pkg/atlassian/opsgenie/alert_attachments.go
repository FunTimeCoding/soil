package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) AlertAttachments(identifier string) []alert.ListedAttachment {
	r, e := c.userClient.Alert.ListAlertsAttachments(
		c.context,
		&alert.ListAttachmentsRequest{
			IdentifierType:  alert.ALERTID,
			IdentifierValue: identifier,
		},
	)
	errors.PanicOnError(e)

	return r.Attachment
}
