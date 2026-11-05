package service

func (s *Service) ReconfigureDnsmasq() error {
	return s.attempt(s.client.ReconfigureDnsmasq)
}
