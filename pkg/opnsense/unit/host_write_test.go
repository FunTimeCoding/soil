package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/opnsense/request"
	"github.com/funtimecoding/soil/pkg/opnsense/unit/opnsense_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestDeleteHostReportsMissingEntry(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithTLS(),
		upstream_server.WithAnswer(`{"result":"not found"}`),
	)
	e := opnsense_tester.New(t, s.URL).DeleteHost(
		"00000000-0000-0000-0000-000000000000",
	)
	assert.True(t, not_found.Is(e))
	assert.StringContains(t, "dnsmasq host not found", e.Error())
}

func TestDeleteHostAcceptsDeleted(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithTLS(),
		upstream_server.WithAnswer(`{"result":"deleted"}`),
	)
	assert.Nil(t, opnsense_tester.New(t, s.URL).DeleteHost("abc"))
}

func TestAddHostRejectionIsValidation(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithTLS(),
		upstream_server.WithAnswer(
			`{"result":"failed","validations":{"host.ip":"bad address"}}`,
		),
	)
	_, e := opnsense_tester.New(t, s.URL).AddHost(request.New())
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "host.ip: bad address", e.Error())
}

func TestAddHostFailingWithoutValidationsIsFault(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithTLS(),
		upstream_server.WithAnswer(`{"result":"failed"}`),
	)
	_, e := opnsense_tester.New(t, s.URL).AddHost(request.New())
	assert.True(t, unexpected.Is(e))
	assert.String(
		t,
		"dnsmasq host save answered failed without validations",
		e.Error(),
	)
}

func TestSetHostOnUnknownEntryIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithTLS(),
		upstream_server.WithAnswer(`{"result":"failed"}`),
	)
	e := opnsense_tester.New(t, s.URL).SetHost(
		"00000000-0000-0000-0000-000000000000",
		request.New(),
	)
	assert.String(
		t,
		"dnsmasq host not found: 00000000-0000-0000-0000-000000000000",
		e.Error(),
	)
}

func TestSetHostRejectionIsValidation(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithTLS(),
		upstream_server.WithAnswer(
			`{"result":"failed","validations":{"host.ip":"bad address"}}`,
		),
	)
	e := opnsense_tester.New(t, s.URL).SetHost("abc", request.New())
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "host.ip: bad address", e.Error())
}

func TestAddHostReturnsIdentifier(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithTLS(),
		upstream_server.WithAnswer(`{"result":"saved","uuid":"a-b-c"}`),
	)
	result, e := opnsense_tester.New(t, s.URL).AddHost(request.New())
	assert.Nil(t, e)
	assert.String(t, "a-b-c", result)
}
