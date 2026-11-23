package graph_tester

import (
	"golang.org/x/oauth2"
	"time"
)

func Valid() *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  "stored-access",
		RefreshToken: "stored-refresh",
		Expiry:       time.Now().Add(time.Hour),
	}
}
