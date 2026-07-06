package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/anonymous_struct"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/call_format"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/defer_close"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/file_identity"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/forbidden_call"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/forbidden_import"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/naming"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/string_concatenation"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/string_constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/struct_literal"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/type_receiver"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unchecked_print_write"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/value_return"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"golang.org/x/tools/go/packages"
)

func check(
	p *packages.Package,
	results *output.Results,
) {
	naming.Check(p, results)
	forbidden_call.Check(p, results)
	forbidden_import.Check(p, results)
	string_concatenation.Check(p, results)
	string_constant.Check(p, results)
	struct_literal.Check(p, results)
	call_format.Check(p, results)
	defer_close.Check(p, results)
	file_identity.Check(p, results)
	type_receiver.Check(p, results)
	unchecked_print_write.Check(p, results)
	anonymous_struct.Check(p, results)
	value_return.Check(p, results)
}
