package service

func (s *Service) PostReactions(identifier string) ([]string, error) {
	return guarded(
		s.guard,
		func() ([]string, error) {
			p, e := s.mattermost.FindPost(identifier)

			if e != nil {
				return nil, e
			}

			result := []string{}

			if !p.HasReactions {
				return result, nil
			}

			reactions, f := s.mattermost.Reactions(p)

			if f != nil {
				return nil, f
			}

			for _, v := range reactions {
				result = append(result, v.EmojiName)
			}

			return result, nil
		},
	)
}
