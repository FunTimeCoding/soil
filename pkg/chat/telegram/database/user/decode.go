package user

import "github.com/funtimecoding/soil/pkg/notation"

func (u *User) Decode(b []byte) {
	notation.MustDecodeBytes(b, u, false)
}
