package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"io"
	"k8s.io/api/core/v1"
	"time"
)

func PodLogsString(
	x context.Context,
	c *cluster.Cluster,
	pod string,
	namespace string,
	container string,
	tail int,
	since string,
	previous bool,
	timestamps bool,
) (string, error) {
	options := &v1.PodLogOptions{
		Follow:    false,
		Previous:  previous,
		Container: container,
	}

	if timestamps {
		options.Timestamps = true
	}

	if tail > 0 {
		lines := int64(tail)
		options.TailLines = &lines
	}

	if since != "" {
		d, e := time.ParseDuration(since)

		if e != nil {
			return "", fmt.Errorf("invalid since duration: %s", e)
		}

		seconds := int64(d.Seconds())
		options.SinceSeconds = &seconds
	}

	stream, e := c.Clientset().CoreV1().Pods(namespace).GetLogs(
		pod,
		options,
	).Stream(x)

	if e != nil {
		return "", e
	}

	defer errors.LogClose(stream)
	b, f := io.ReadAll(stream)

	if f != nil {
		return "", f
	}

	return string(b), nil
}
