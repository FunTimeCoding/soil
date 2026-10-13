package board_entry

import "time"

func New(
	instance string,
	project string,
	projectIdentifier int64,
	projectLink string,
	reference string,
	status string,
	identifier int64,
	link string,
	updated time.Time,
) *Entry {
	return &Entry{
		Instance:          instance,
		Project:           project,
		ProjectIdentifier: projectIdentifier,
		ProjectLink:       projectLink,
		Reference:         reference,
		Status:            status,
		Identifier:        identifier,
		Link:              link,
		Updated:           updated,
	}
}
