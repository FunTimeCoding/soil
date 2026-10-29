package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/mock_client/page_entry"

func New() *Client {
	return &Client{pages: map[string]*page_entry.Entry{}, nextIdentifier: 1}
}
