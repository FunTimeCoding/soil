package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"testing"
)

func TestBareMarkerSuppressesEveryKey(t *testing.T) {
	assert.True(
		t,
		suppress.CommentSuppresses("// goanalyze:ignore", "forbidden_call"),
	)
	assert.True(
		t,
		suppress.CommentSuppresses("// goanalyze:ignore", "restricted_literal"),
	)
}

func TestKeyedMarkerSuppressesOnlyItsKey(t *testing.T) {
	assert.True(
		t,
		suppress.CommentSuppresses(
			"// goanalyze:ignore forbidden_call",
			"forbidden_call",
		),
	)
	assert.False(
		t,
		suppress.CommentSuppresses(
			"// goanalyze:ignore forbidden_call",
			"restricted_call",
		),
	)
}

func TestMarkerWithoutSpaceAfterSlashesSuppresses(t *testing.T) {
	assert.True(
		t,
		suppress.CommentSuppresses(
			"//goanalyze:ignore forbidden_call",
			"forbidden_call",
		),
	)
}

func TestTrailingWhitespaceDoesNotChangeKey(t *testing.T) {
	assert.True(
		t,
		suppress.CommentSuppresses(
			"// goanalyze:ignore forbidden_call  \t",
			"forbidden_call",
		),
	)
}

func TestMarkerAfterSecurityDirectiveSuppresses(t *testing.T) {
	assert.True(
		t,
		suppress.CommentSuppresses(
			"// #nosec G204 goanalyze:ignore forbidden_call",
			"forbidden_call",
		),
	)
	assert.False(
		t,
		suppress.CommentSuppresses(
			"// #nosec G204 goanalyze:ignore forbidden_call",
			"restricted_call",
		),
	)
}

func TestBareMarkerAfterSecurityDirectiveSuppressesEveryKey(t *testing.T) {
	assert.True(
		t,
		suppress.CommentSuppresses(
			"// #nosec G204 goanalyze:ignore",
			"forbidden_call",
		),
	)
}

func TestSecurityDirectiveAloneSuppressesNothing(t *testing.T) {
	assert.False(
		t,
		suppress.CommentSuppresses("// #nosec G204", "forbidden_call"),
	)
}

func TestOneOfSeveralKeyedMarkersSuppresses(t *testing.T) {
	text := "// goanalyze:ignore forbidden_call goanalyze:ignore restricted_literal"
	assert.True(t, suppress.CommentSuppresses(text, "forbidden_call"))
	assert.True(t, suppress.CommentSuppresses(text, "restricted_literal"))
	assert.False(t, suppress.CommentSuppresses(text, "struct_literal"))
}

func TestMarkerFollowedByProseSuppressesEveryKey(t *testing.T) {
	assert.True(
		t,
		suppress.CommentSuppresses(
			"// goanalyze:ignore Because the runner owns the process",
			"forbidden_call",
		),
	)
}

func TestMarkerInsideWordDoesNotSuppress(t *testing.T) {
	assert.False(
		t,
		suppress.CommentSuppresses(
			"// see goanalyze:ignored for details",
			"forbidden_call",
		),
	)
	assert.False(
		t,
		suppress.CommentSuppresses(
			"// xgoanalyze:ignore forbidden_call",
			"forbidden_call",
		),
	)
}

func TestKeyMustMatchExactly(t *testing.T) {
	assert.False(
		t,
		suppress.CommentSuppresses(
			"// goanalyze:ignore forbidden_call_extra",
			"forbidden_call",
		),
	)
	assert.False(
		t,
		suppress.CommentSuppresses(
			"// goanalyze:ignore forbidden",
			"forbidden_call",
		),
	)
}

func TestBlockCommentDoesNotSuppress(t *testing.T) {
	assert.False(
		t,
		suppress.CommentSuppresses(
			"/* goanalyze:ignore forbidden_call */",
			"forbidden_call",
		),
	)
}
