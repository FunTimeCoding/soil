package service

type RegisterClientResponse struct {
	ClientIdentifier string
	ClientSecret     string
	RedirectLocators []string
}
