package option

import "github.com/funtimecoding/soil/pkg/directory/face"

type Directory struct {
	Address          string
	ServiceTokens    []string
	Issuer           string
	ClientIdentifier string
	ClientSecret     string
	EncryptionSecret string
	PublicLocator    string
	Host             string
	Source           func() face.Client
}
