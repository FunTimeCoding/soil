package summary

import "github.com/funtimecoding/soil/pkg/backrest/response"

type Summary struct {
	Stanza string
	Status int
	Latest map[string]*response.Backup
}
