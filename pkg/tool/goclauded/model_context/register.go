package model_context

func (s *Server) register() {
	s.registerLifecycle()
	s.registerSession()
	s.registerRoster()
	s.registerHistory()
	s.registerMessaging()
	s.registerConversations()
	s.registerStatus()
}
