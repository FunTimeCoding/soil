package option

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board"

type Dashboard struct {
	Address          string
	Version          string
	Board            *board.Board
	PostgresLocator  string
	LitePath         string
	Issuer           string
	ClientIdentifier string
	ClientSecret     string
	EncryptionSecret string
	PublicLocator    string
}
