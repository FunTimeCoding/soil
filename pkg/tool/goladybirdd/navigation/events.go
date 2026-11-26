package navigation

import browserConstant "github.com/funtimecoding/soil/pkg/browser/constant"

func Events() []string {
	return []string{
		browserConstant.NavigationStartedEvent,
		browserConstant.LoadEvent,
		browserConstant.NavigationFailedEvent,
		browserConstant.NavigationAbortedEvent,
		browserConstant.FragmentNavigatedEvent,
		browserConstant.HistoryUpdatedEvent,
	}
}
