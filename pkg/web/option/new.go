package option

func New(
	address string,
	tokens ...string,
) *Web {
	return &Web{Address: address, ServiceTokens: tokens}
}
