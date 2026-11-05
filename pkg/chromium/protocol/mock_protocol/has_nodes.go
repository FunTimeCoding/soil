package mock_protocol

func (p *Protocol) HasNodes(_ string) (bool, error) {
	return false, nil
}
