package document_cache

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"hash/maphash"
)

type Cache struct {
	seed     maphash.Seed
	current  map[string]*file_report.Report
	previous map[string]*file_report.Report
}
