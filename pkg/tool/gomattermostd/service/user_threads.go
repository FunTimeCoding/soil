package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) UserThreads(o model.GetUserThreadsOpts) (*model.Threads, error) {
	return guarded(
		s.guard,
		func() (*model.Threads, error) {
			me, e := s.mattermost.Me()

			if e != nil {
				return nil, e
			}

			return s.mattermost.UserThreads(me, s.mattermost.DefaultTeam(), o)
		},
	)
}
