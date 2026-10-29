package option

import "github.com/funtimecoding/soil/pkg/directory/face"

type Option struct {
	Address               string
	PostgresLocator       string
	LitePath              string
	Secret                string
	SuperUserMail         string
	SuperUserPassword     string
	Issuer                string
	AdminClientIdentifier string
	AdminClientSecret     string
	ServiceTokens         []string
	Directory             func() face.Client
}
