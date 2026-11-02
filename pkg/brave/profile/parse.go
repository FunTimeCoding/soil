package profile

import (
	"encoding/base64"
	"github.com/funtimecoding/soil/pkg/brave/helper"
	"path/filepath"
)

func Parse(
	base string,
	profile string,
	s string,
) (*Profile, error) {
	b, e := base64.StdEncoding.DecodeString(s)

	if e != nil {
		return nil, e
	}

	_, mail, _, _ := helper.ParseAccounts(b)

	return &Profile{
		Profile: profile,
		Path:    filepath.Join(base, profile),
		Email:   mail,
	}, nil
}
