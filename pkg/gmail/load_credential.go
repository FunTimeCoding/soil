package gmail

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gmail/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func (c *Client) loadCredential() *oauth2.Config {
	result, e := google.ConfigFromJSON(
		system.ReadBytes(c.directory, constant.CredentialFile),
		gmail.GmailReadonlyScope,
	)
	errors.PanicOnError(e)

	return result
}
