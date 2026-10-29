package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/directory/types/entry"
	"github.com/funtimecoding/soil/pkg/directory/types/search_record"
)

type Client interface {
	Add(
		distinguishedName string,
		attributes map[string][]string,
	) error
	Authenticate(
		account string,
		password string,
	) (*entry.Entry, error)
	Base() string
	Delete(distinguishedName string) error
	InGroup(account string) (bool, error)
	Modify(
		distinguishedName string,
		attributes map[string][]string,
	) error
	Reachable(x context.Context) error
	Search(
		filter string,
		attributes []string,
	) ([]*search_record.Record, error)
	SetPassword(
		distinguishedName string,
		password string,
	) error
}
