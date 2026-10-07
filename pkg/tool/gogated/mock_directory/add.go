package mock_directory

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/directory/types/entry"
)

func (d *Directory) Add(
	unique string,
	account string,
	mail string,
	password string,
	member bool,
) {
	d.entries = append(
		d.entries,
		entry.New(
			unique,
			account,
			mail,
			account,
			fmt.Sprintf("uid=%s,ou=people", account),
		),
	)
	d.passwords[account] = password
	d.members[account] = member
}
