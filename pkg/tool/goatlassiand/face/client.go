package face

import "context"

type Client interface {
	ChecklistClient
	PageClient
	IssueClient
	Address() string
	Reachable(x context.Context) error
	Close()
}
