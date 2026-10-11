package package_server

func New(signatureKey string) *Server {
	return &Server{signatureKey: signatureKey}
}
