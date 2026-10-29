package service

func (s *Service) ReconfigureDnsmasq() error {
	return s.source.ReconfigureDnsmasq()
}
