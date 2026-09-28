package fosite_client

import "strings"

func (c *Client) GetRedirectURIs() []string {
	return strings.Split(c.Row.RedirectLocators, "\n")
}
