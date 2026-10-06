package bearer

func New(token string) *Bearer {
	return &Bearer{token: token}
}
