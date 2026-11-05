package mock_protocol

func (p *Protocol) Outer(_ string) (string, error) {
	return "", nil
}
