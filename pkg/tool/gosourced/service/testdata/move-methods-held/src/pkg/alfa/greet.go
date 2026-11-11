package alfa

import "strings"

func (s *Service) Greet(who string) string {
	return strings.ToUpper(who)
}
