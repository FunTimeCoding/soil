package document_cache

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"hash/maphash"
)

func New() *Cache {
	return &Cache{
		seed:     maphash.MakeSeed(),
		current:  make(map[string]*file_report.Report),
		previous: make(map[string]*file_report.Report),
	}
}
