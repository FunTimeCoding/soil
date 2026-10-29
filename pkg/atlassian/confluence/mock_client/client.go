package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/mock_client/page_entry"

type Client struct {
	pages          map[string]*page_entry.Entry
	nextIdentifier int
	Closed         int
}
