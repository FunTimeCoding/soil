package fixture

func (u *User) UnknownField() map[string]any {
	return u.Unknown
}
