package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/page_entry"

func New() *Client {
	return &Client{pages: map[string]*page_entry.Entry{}, nextIdentifier: 1}
}
