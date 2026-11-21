package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/subject"
	"testing"
)

func TestSubjectKindReadsEveryForm(t *testing.T) {
	for _, v := range [][2]string{
		{"CVE-2099-00001", "vulnerability"},
		{"upgrade:bookworm:victor", "upgrade"},
		{"upgrade:bookworm:linux (backports)", "upgrade"},
		{"release:juliett", "release"},
		{"reboot:foxtrot", "reboot"},
		{"kernel:echo", "kernel"},
		{"kernel:", ""},
		{"upgrade:bookworm", ""},
		{"upgrade::victor", ""},
		{"upgrade:bookworm:", ""},
		{"release:", ""},
		{"bogus:alfa", ""},
		{"juliett", ""},
	} {
		assert.String(t, v[1], subject.Kind(v[0]))
	}

	assert.True(t, validation.Is(subject.Check("juliett")))
	assert.FatalOnError(t, subject.Check("reboot:foxtrot"))
	assert.String(
		t,
		"upgrade:bookworm:victor",
		subject.Upgrade("bookworm", "victor"),
	)
	assert.String(
		t,
		"release:golf",
		subject.Host(constant.RemediationRelease, "golf"),
	)
}
