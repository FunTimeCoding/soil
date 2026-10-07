package mock_directory

import (
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/directory/types/entry"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (d *Directory) Authenticate(
	account string,
	password string,
) (*entry.Entry, error) {
	if d.failure != nil {
		return nil, d.failure
	}

	for _, n := range d.entries {
		if n.Account != account && n.Mail != account {
			continue
		}

		if d.passwords[n.Account] != password {
			return nil, validation.New("invalid credentials")
		}

		return n, nil
	}

	return nil, not_found.New(constant.Subject, account)
}
