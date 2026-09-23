package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/opnsense/request"
	"testing"
)

func TestDeleteHostReportsMissingEntry(t *testing.T) {
	s := canned(`{"result":"not found"}`)
	defer s.Close()
	e := client(s).DeleteHost("00000000-0000-0000-0000-000000000000")
	assert.True(t, not_found.Is(e))
	assert.StringContains(t, "dnsmasq host not found", e.Error())
}

func TestDeleteHostAcceptsDeleted(t *testing.T) {
	s := canned(`{"result":"deleted"}`)
	defer s.Close()
	assert.Nil(t, client(s).DeleteHost("abc"))
}

func TestAddHostRejectionIsValidation(t *testing.T) {
	s := canned(`{"result":"failed","validations":{"host.ip":"bad address"}}`)
	defer s.Close()
	_, e := client(s).AddHost(request.New())
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "host.ip: bad address", e.Error())
}

func TestAddHostReturnsIdentifier(t *testing.T) {
	s := canned(`{"result":"saved","uuid":"a-b-c"}`)
	defer s.Close()
	result, e := client(s).AddHost(request.New())
	assert.Nil(t, e)
	assert.String(t, "a-b-c", result)
}
