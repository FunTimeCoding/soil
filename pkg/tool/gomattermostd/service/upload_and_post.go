package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) UploadAndPost(
	h *model.Channel,
	name string,
	content []byte,
	message string,
) (*model.Post, *model.FileInfo, error) {
	var file *model.FileInfo
	p, e := guarded(
		s.guard,
		func() (*model.Post, error) {
			var f error
			file, f = s.mattermost.UploadFile(h, name, content)

			if f != nil {
				return nil, f
			}

			return s.mattermost.Post(
				&model.Post{
					ChannelId: h.Id,
					Message:   message,
					FileIds:   []string{file.Id},
				},
			)
		},
	)

	if e != nil {
		return nil, nil, e
	}

	return p, file, nil
}
