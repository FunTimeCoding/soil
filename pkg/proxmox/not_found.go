package proxmox

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/luthermonson/go-proxmox"
)

func notFound(
	e error,
	kind string,
	identifier any,
) error {
	if errors.Is(e, proxmox.ErrNotFound) {
		return not_found.New(kind, identifier)
	}

	return e
}
