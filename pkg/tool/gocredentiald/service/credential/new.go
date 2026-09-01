package credential

import "time"

func New(
	identifier string,
	path string,
	title string,
	user string,
	locator string,
	modifiedAt time.Time,
) *Credential {
	return &Credential{
		Identifier: identifier,
		Path:       path,
		Title:      title,
		User:       user,
		Locator:    locator,
		ModifiedAt: modifiedAt,
	}
}
