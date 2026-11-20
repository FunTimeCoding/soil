package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"slices"
	"time"
)

func (c *Client) PutSystem(
	_ context.Context,
	name string,
	b client.SystemSnapshot,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return c.failure
	}

	if !constant.SystemNamePattern.MatchString(name) {
		return validation.New("system name malformed: %s", name)
	}

	if b.Release == "" {
		return validation.New("system release missing: %s", name)
	}

	tracked := slices.Contains(constant.SupportedReleases, b.Release) ||
		slices.Contains(constant.ExtendedReleases, b.Release)

	if len(b.Packages) == 0 && tracked {
		return validation.New(
			"system without packages on a tracked release: %s",
			name,
		)
	}

	for _, p := range b.Packages {
		if p == nil || p.Source == "" || p.SourceVersion == "" {
			return validation.New(
				"package without source name or version on system: %s",
				name,
			)
		}
	}

	c.snapshots[name] = b
	c.updated[name] = time.Now()
	delete(c.details, name)

	return nil
}
