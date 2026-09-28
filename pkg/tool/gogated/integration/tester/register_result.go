package tester

type RegisterResult struct {
	ClientIdentifier string   `json:"client_id"`
	ClientSecret     string   `json:"client_secret"`
	RedirectURIs     []string `json:"redirect_uris"`
}
