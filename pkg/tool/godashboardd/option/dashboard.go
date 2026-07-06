package option

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board"

type Dashboard struct {
	Port             int
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
