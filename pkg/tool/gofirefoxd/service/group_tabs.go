package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/firefox/constant"
)

func (s *Service) GroupTabs(
	tabIdentifiers []int,
	groupIdentifier int,
	title string,
	color string,
) (int, error) {
	if len(tabIdentifiers) == 0 {
		return 0, validation.New(constant.TabsRequired)
	}

	return s.client.GroupTabs(tabIdentifiers, groupIdentifier, title, color)
}
