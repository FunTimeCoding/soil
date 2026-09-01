package credential

import "time"

type Credential struct {
	Identifier string
	Path       string
	Title      string
	User       string
	Locator    string
	ModifiedAt time.Time
}
