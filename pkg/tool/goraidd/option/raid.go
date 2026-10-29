package option

import "github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"

type Raid struct {
	Address          string
	PostgresLocator  string
	LitePath         string
	LogCachePath     string
	ElitePath        string
	OutputPath       string
	Issuer           string
	ClientIdentifier string
	ClientSecret     string
	EncryptionSecret string
	PublicLocator    string
	ServiceTokens    []string
	Parser           func() face.Client
}
