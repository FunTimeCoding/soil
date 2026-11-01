package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (s *Service) recordIntent(
	instance string,
	operation string,
) {
	if s.netbox == nil {
		return
	}

	status, known := constant.IntentStatus[operation]

	if !known {
		return
	}

	e := s.netboxGuard.Attempt(
		context.Background(),
		func(x context.Context) error {
			body := client.UpdateDeviceRequest{Status: &status}
			_, f := s.netbox.UpdateDevice(x, instance, body)

			return f
		},
	)

	if e == nil {
		return
	}

	if classify.Reportable(e) {
		s.reporter.CaptureException(e)

		return
	}

	s.logger.Plain("%s status not recorded: %v", instance, e)
}
