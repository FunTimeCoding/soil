package package_server

func New(
	token string,
	signatureKey string,
) *Server {
	return &Server{token: token, signatureKey: signatureKey}
}
