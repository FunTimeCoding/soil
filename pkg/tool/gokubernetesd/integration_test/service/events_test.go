package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"testing"
)

func TestEvents(t *testing.T) {
	s := service_tester.New(t)
	s.AddEvent("default", "Pulled", "pulled image", "Normal", "Pod", "nginx")
	result, e := s.Service.Events(
		context.Background(),
		"test",
		service.EventsQuery{
			Namespace: "default",
			Limit:     50,
		},
	)
	assert.Nil(t, e)
	assert.Count(t, 1, result)
	assert.String(t, "Pulled", result[0].Reason)
}

func TestEventsMutedFiltered(t *testing.T) {
	s := service_tester.New(t)
	s.AddEvent(
		"kube-system",
		"DNSConfigForming",
		"nameserver limits",
		"Warning",
		"Pod",
		"calico",
	)
	s.AddEvent(
		"kube-system",
		"Pulled",
		"pulled image",
		"Normal",
		"Pod",
		"coredns",
	)
	s.Mute("DNSConfigForming", "", "")
	result, e := s.Service.Events(
		context.Background(),
		"test",
		service.EventsQuery{
			Namespace: "kube-system",
			Limit:     50,
		},
	)
	assert.Nil(t, e)
	assert.Count(t, 1, result)
	assert.String(t, "Pulled", result[0].Reason)
}

func TestEventsMutedUnfiltered(t *testing.T) {
	s := service_tester.New(t)
	s.AddEvent(
		"kube-system",
		"DNSConfigForming",
		"nameserver limits",
		"Warning",
		"Pod",
		"calico",
	)
	s.AddEvent(
		"kube-system",
		"Pulled",
		"pulled image",
		"Normal",
		"Pod",
		"coredns",
	)
	s.Mute("DNSConfigForming", "", "")
	result, e := s.Service.Events(
		context.Background(),
		"test",
		service.EventsQuery{
			Namespace:  "kube-system",
			Limit:      50,
			Unfiltered: true,
		},
	)
	assert.Nil(t, e)
	assert.Count(t, 2, result)
}
