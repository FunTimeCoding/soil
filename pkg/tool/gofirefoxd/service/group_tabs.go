package service

import (
	browserConstant "github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (s *Service) GroupTabs(
	tabIdentifiers []int,
	groupIdentifier int,
	title string,
	color string,
) (int, error) {
	if len(tabIdentifiers) == 0 {
		return 0, validation.New(browserConstant.FirefoxTabsRequired)
	}

	return s.client.GroupTabs(tabIdentifiers, groupIdentifier, title, color)
}
