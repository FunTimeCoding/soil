package service

import (
	"github.com/funtimecoding/soil/pkg/keepass"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/tobischo/gokeepasslib/v3"
)

func applyFields(
	entry *gokeepasslib.Entry,
	fields map[string]string,
) {
	for key, content := range fields {
		protected := key == constant.PasswordKey
		found := false

		for i := range entry.Values {
			if entry.Values[i].Key != key {
				continue
			}

			entry.Values[i].Value.Content = content
			found = true
		}

		if found {
			continue
		}

		entry.Values = append(
			entry.Values,
			keepass.NewValue(key, content, protected),
		)
	}
}
