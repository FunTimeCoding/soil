package store

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/user"

func (s *Store) RefreshDirectoryUser(
	row *user.User,
	account string,
	mail string,
	name string,
) error {
	if row.Account == account && row.Mail == mail && row.Name == name {
		return nil
	}

	row.Account = account
	row.Mail = mail
	row.Name = name

	return s.mapper.Save(row).Error
}
