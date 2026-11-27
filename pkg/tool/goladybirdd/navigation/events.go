package navigation

import browserConstant "github.com/funtimecoding/soil/pkg/browser/constant"

func Events() []string {
	return []string{
		browserConstant.LadybirdNavigationStartedEvent,
		browserConstant.LadybirdLoadEvent,
		browserConstant.LadybirdNavigationFailedEvent,
		browserConstant.LadybirdNavigationAbortedEvent,
		browserConstant.LadybirdFragmentNavigatedEvent,
		browserConstant.LadybirdHistoryUpdatedEvent,
	}
}
