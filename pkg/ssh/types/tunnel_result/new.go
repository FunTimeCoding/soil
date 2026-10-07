package tunnel_result

func New(localPort int) *Result {
	return &Result{LocalPort: localPort}
}
