package login_session

func New(
	identifier string,
	authorizeRequest string,
) *LoginSession {
	return &LoginSession{
		Identifier:       identifier,
		AuthorizeRequest: authorizeRequest,
	}
}
