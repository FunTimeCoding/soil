package locator

func NewEndpoint(
	host string,
	port int,
	secure bool,
) *Locator {
	result := New(host).Port(port)

	if !secure {
		result.Insecure()
	}

	return result
}
