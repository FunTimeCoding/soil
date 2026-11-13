package pointer_cache

import (
	"github.com/funtimecoding/soil/pkg/lint/types/pointer_entry"
	"hash/maphash"
)

type Cache struct {
	seed     maphash.Seed
	entries  map[string]*pointer_entry.Entry
	globals  map[string]string
	checks   int
	auditing bool
	Mismatch func(string, string, string, string)
}
