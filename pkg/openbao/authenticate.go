package openbao

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/openbao/constant"
	"github.com/openbao/openbao/api/auth/kubernetes/v2"
	"time"
)

func (c *Client) authenticate(x context.Context) error {
	if c.role == "" {
		return nil
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	if time.Now().Before(c.expiry) {
		return nil
	}

	a, e := kubernetes.NewKubernetesAuth(
		c.role,
		kubernetes.WithServiceAccountTokenPath(c.tokenPath),
	)

	if e != nil {
		return fmt.Errorf("prepare login: %w", e)
	}

	s, f := c.client.Auth().Login(x, a)

	if f != nil {
		return fmt.Errorf("log in: %w", f)
	}

	lease := time.Duration(s.Auth.LeaseDuration) * time.Second
	c.expiry = time.Now().Add(lease - lease/constant.RenewDenominator)

	return nil
}
