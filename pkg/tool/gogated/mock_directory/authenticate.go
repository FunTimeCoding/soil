package mock_directory

import (
	"github.com/funtimecoding/soil/pkg/directory"
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (d *Directory) Authenticate(
	account string,
	password string,
) (*directory.Entry, error) {
	if d.failure != nil {
		return nil, d.failure
	}

	for _, entry := range d.entries {
		if entry.Account != account && entry.Mail != account {
			continue
		}

		if d.passwords[entry.Account] != password {
			return nil, validation.New("invalid credentials")
		}

		return entry, nil
	}

	return nil, not_found.New(constant.Subject, account)
}
