package mock_token_store

import (
	"golang.org/x/oauth2"
	"testing"
)

func New(
	t *testing.T,
	token *oauth2.Token,
) *Store {
	t.Helper()

	return &Store{token: token}
}
