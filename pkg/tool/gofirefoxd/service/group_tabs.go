package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	firefoxConstant "github.com/funtimecoding/soil/pkg/firefox/constant"
)

func (s *Service) GroupTabs(
	tabIdentifiers []int,
	groupIdentifier int,
	title string,
	color string,
) (int, error) {
	if len(tabIdentifiers) == 0 {
		return 0, validation.New(firefoxConstant.TabsRequired)
	}

	return s.client.GroupTabs(tabIdentifiers, groupIdentifier, title, color)
}
