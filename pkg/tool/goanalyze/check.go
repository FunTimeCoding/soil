package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/anonymous_struct"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/call_format"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/constant_declaration"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/defer_close"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/error_wrap_verb"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/expected_first"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/file_identity"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/forbidden_call"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/forbidden_import"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/naming"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/omit_empty_zero"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/restricted_call"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/stray_comment"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/stray_constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/stray_variable"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/string_concatenation"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/string_constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/struct_literal"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/type_receiver"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unchecked_print_write"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/value_return"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"golang.org/x/tools/go/packages"
)

func check(
	p *packages.Package,
	results *output.Results,
	comment bool,
	faces *face.Set,
) {
	naming.Check(p, results, faces)
	forbidden_call.Check(p, results)
	restricted_call.Check(p, results)
	forbidden_import.Check(p, results)
	string_concatenation.Check(p, results)
	string_constant.Check(p, results)
	expected_first.Check(p, results)
	struct_literal.Check(p, results)
	call_format.Check(p, results)
	defer_close.Check(p, results)
	error_wrap_verb.Check(p, results)
	omit_empty_zero.Check(p, results)
	file_identity.Check(p, results)
	type_receiver.Check(p, results)
	unchecked_print_write.Check(p, results)
	anonymous_struct.Check(p, results)
	value_return.Check(p, results)
	stray_variable.Check(p, results)
	stray_constant.Check(p, results)
	constant_declaration.Check(p, results)

	if comment {
		stray_comment.Check(p, results)
	}
}
