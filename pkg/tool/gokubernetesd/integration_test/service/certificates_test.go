package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester"
	"testing"
	"time"
)

func TestCertificates(t *testing.T) {
	s := service_tester.New(t)
	s.AddCertificate(
		"gate",
		"gate-tls",
		[]string{"gate.s3n.sh"},
		time.Now().Add(89*24*time.Hour),
		true,
	)
	s.AddCertificate(
		"sentry",
		"sentry-tls",
		[]string{"sentry.s3n.sh"},
		time.Now().Add(28*24*time.Hour),
		true,
	)
	result, e := s.Service.Certificates(context.Background(), "test")
	assert.Nil(t, e)
	assert.Count(t, 2, result)
	assert.String(t, "sentry-tls", result[0].Name)
	assert.String(t, "gate-tls", result[1].Name)
}

func TestCertificatesSortedByExpiry(t *testing.T) {
	s := service_tester.New(t)
	s.AddCertificate(
		"a",
		"long-lived",
		[]string{"a.example"},
		time.Now().Add(90*24*time.Hour),
		true,
	)
	s.AddCertificate(
		"b",
		"expiring-soon",
		[]string{"b.example"},
		time.Now().Add(7*24*time.Hour),
		true,
	)
	s.AddCertificate(
		"c",
		"medium",
		[]string{"c.example"},
		time.Now().Add(30*24*time.Hour),
		true,
	)
	result, e := s.Service.Certificates(context.Background(), "test")
	assert.Nil(t, e)
	assert.String(t, "expiring-soon", result[0].Name)
	assert.String(t, "medium", result[1].Name)
	assert.String(t, "long-lived", result[2].Name)
}
