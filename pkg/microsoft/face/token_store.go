package face

import "golang.org/x/oauth2"

type TokenStore interface {
	FindToken() (*oauth2.Token, bool, error)
	SaveToken(t *oauth2.Token) error
}
