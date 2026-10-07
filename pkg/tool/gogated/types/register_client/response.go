package register_client

type Response struct {
	ClientIdentifier string
	ClientSecret     string
	RedirectLocators []string
}
