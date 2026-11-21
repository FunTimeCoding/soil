package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	"testing"
)

func TestEarlierFlagColumnBecomesSubject(t *testing.T) {
	s, m := security_tester.New(t).EarlierFlags("upgrade:bookworm:victor")
	f, e := s.BySubject("upgrade:bookworm:victor")
	assert.FatalOnError(t, e)
	assert.String(t, "kept", f.Note)
	migrator := m.Migrator()
	assert.True(t, migrator.HasColumn(flag.Stub(), constant.FlagSubjectColumn))
	assert.False(t, migrator.HasColumn(flag.Stub(), constant.FlagEarlierColumn))
	assert.False(t, migrator.HasIndex(flag.Stub(), constant.FlagEarlierIndex))
	assert.True(t, migrator.HasIndex(flag.Stub(), "idx_security_flags_subject"))
}
