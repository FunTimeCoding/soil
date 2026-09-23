package fixture

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (u *User) UnmarshalJSON(b []byte) error {
	type Alias User
	v := (*Alias)(u)

	return notation.UnmarshalUnknown(b, v, constant.UnknownField)
}
