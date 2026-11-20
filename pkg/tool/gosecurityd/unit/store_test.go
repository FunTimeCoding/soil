package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
)

func TestCreatingFlaggedVulnerabilityAgainIsConflict(t *testing.T) {
	s := security_tester.New(t).Store()
	first := flag.Stub()
	first.Vulnerability = vulnerabilityConstant.KnownIdentifier
	assert.FatalOnError(t, s.Create(first))
	second := flag.Stub()
	second.Vulnerability = vulnerabilityConstant.KnownIdentifier
	e := s.Create(second)
	assert.True(t, conflict.Is(e))
	assert.String(t, "flag already exists: CVE-2024-3094", e.Error())
}

func TestFlagRacingAnotherSubmitIsConflict(t *testing.T) {
	o := security_tester.New(t)
	o.NationalDatabase.During(
		func() {
			competing := flag.Stub()
			competing.Vulnerability = vulnerabilityConstant.KnownIdentifier
			assert.FatalOnError(t, o.Store().Create(competing))
		},
	)
	_, e := o.Flag().Flag(
		t.Context(),
		vulnerabilityConstant.KnownIdentifier,
		"",
	)
	assert.True(t, conflict.Is(e))
}
