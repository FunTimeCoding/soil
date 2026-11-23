package graph_tester

import (
	"golang.org/x/oauth2"
	"time"
)

func Expired() *oauth2.Token {
	result := Valid()
	result.Expiry = time.Now().Add(-time.Hour)

	return result
}
