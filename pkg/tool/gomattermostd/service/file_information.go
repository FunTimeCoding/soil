package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) FileInformation(identifier string) (*model.FileInfo, error) {
	return guarded(
		s.guard,
		func() (*model.FileInfo, error) {
			return s.mattermost.FileInformation(identifier)
		},
	)
}
