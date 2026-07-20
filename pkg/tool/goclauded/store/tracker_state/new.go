package tracker_state

import "time"

func New(
	identifier string,
	payload string,
	updatedAt time.Time,
) *Record {
	return &Record{
		Identifier: identifier,
		Payload:    payload,
		UpdatedAt:  updatedAt,
	}
}
