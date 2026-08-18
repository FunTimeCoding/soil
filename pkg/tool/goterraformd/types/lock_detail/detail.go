package lock_detail

import "time"

type Detail struct {
	Identifier string    `json:"ID"`
	Operation  string    `json:"Operation"`
	Who        string    `json:"Who"`
	Version    string    `json:"Version"`
	Created    time.Time `json:"Created"`
}
