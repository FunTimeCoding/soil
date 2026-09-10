package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/finding"
)

func (s *Service) Findings() ([]*finding.Finding, error) {
	var result []*finding.Finding
	available, availableError := s.store.AvailableNames()

	if availableError != nil {
		return nil, availableError
	}

	if len(available) == 0 {
		total, e := s.store.CountPoolNames()

		if e != nil {
			return nil, e
		}

		result = append(
			result,
			finding.New(
				constant.PoolExhausted,
				"",
				fmt.Sprintf(
					"no callsign available, all %d pool names held",
					total,
				),
				1,
			),
		)
	}

	missing, missingError := s.missingTranscriptFindings()

	if missingError != nil {
		return nil, missingError
	}

	result = append(result, missing...)
	orphans, orphanError := s.store.OrphanTrackerStates()

	if orphanError != nil {
		return nil, orphanError
	}

	for _, i := range orphans {
		result = append(
			result,
			finding.New(
				constant.OrphanTrackerState,
				i,
				"tracker watermark without a session",
				1,
			),
		)
	}

	queued, queueError := s.store.UnownedQueueCallsigns()

	if queueError != nil {
		return nil, queueError
	}

	queuedRows, queuedCountError := s.store.CountUnownedQueue()

	if queuedCountError != nil {
		return nil, queuedCountError
	}

	if f := unownedFinding(
		constant.UnownedQueue,
		"queue entries",
		queued,
		queuedRows,
	); f != nil {
		result = append(result, f)
	}

	notified, notificationError := s.store.UnownedNotificationCallsigns()

	if notificationError != nil {
		return nil, notificationError
	}

	notifiedRows, notifiedCountError := s.store.CountUnownedNotification()

	if notifiedCountError != nil {
		return nil, notifiedCountError
	}

	if f := unownedFinding(
		constant.UnownedNotification,
		"notifications",
		notified,
		notifiedRows,
	); f != nil {
		result = append(result, f)
	}

	result = append(result, s.migrationFindings()...)
	stale, staleError := s.store.StaleCallsignSessions(
		s.clock().Add(-constant.CallsignReleaseWindow),
	)

	if staleError != nil {
		return nil, staleError
	}

	for _, i := range stale {
		result = append(
			result,
			finding.New(
				constant.StaleCallsign,
				i.CallsignValue(),
				"callsign held past the release window, sweep not running",
				1,
			),
		)
	}

	return result, nil
}
