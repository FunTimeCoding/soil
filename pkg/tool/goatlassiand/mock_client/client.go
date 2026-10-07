package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/page_entry"

type Client struct {
	pages          map[string]*page_entry.Entry
	nextIdentifier int
}
