package user

import "github.com/funtimecoding/soil/pkg/notation"

func (u *User) Encode() []byte {
	return notation.Marshal(u)
}
