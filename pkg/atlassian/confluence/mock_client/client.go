package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/mock_client/page_entry"
	"github.com/funtimecoding/soil/pkg/failure"
)

type Client struct {
	failures       *failure.Failure
	pages          map[string]*page_entry.Entry
	nextIdentifier int
	Failure        error
	Closed         int
}
