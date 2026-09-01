package entry_detail

import "time"

type Detail struct {
	Identifier string
	Path       string
	Title      string
	Fields     map[string]string
	ModifiedAt time.Time
}
