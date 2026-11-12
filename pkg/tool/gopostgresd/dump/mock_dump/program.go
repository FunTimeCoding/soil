package mock_dump

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/dump/mock_dump/restoration"

type Program struct {
	Failure  error
	Tables   int
	Dumped   []string
	Restored []*restoration.Restoration
}
