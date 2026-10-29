package service

func (s *Service) Close() {
	s.jiraGuard.Close()
	s.confluenceGuard.Close()
	s.jira.Close()
	s.confluence.Close()
}
