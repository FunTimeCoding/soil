package gmail

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"golang.org/x/oauth2"
)

func (c *Client) tokenFlow(
	o *oauth2.Config,
	skipExisting bool,
) *oauth2.Token {
	if !skipExisting {
		if t := c.selectExistingToken(); t != nil {
			return t
		}
	}

	link := o.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	if !skipExisting {
		system.OpenBrowser(link)
	} else {
		fmt.Printf("Link: %s\n", link)
	}

	c.callback.Start()
	code := c.callback.WaitForCallback()
	fmt.Printf("Callback code: %s\n", code)

	return c.exchange(o, code)
}
