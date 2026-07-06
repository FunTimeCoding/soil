package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) SavedSearch() *alert.SavedSearchResult {
	result, e := c.userClient.Alert.ListSavedSearches(
		c.context,
		&alert.ListSavedSearchRequest{},
	)
	errors.PanicOnError(e)

	return result
}
