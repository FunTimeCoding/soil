package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (r *Resolver) undeclaredMember(
	reference string,
	directory string,
	head string,
) bool {
	member := pointer.MemberName(reference, directory)

	if member == "" || !r.DeclaresType(directory, head) {
		return false
	}

	return !r.Declares(
		directory,
		join.Empty(head, stringsConstant.Dot, member),
	) && !r.Declares(
		directory,
		join.Empty(head, stringsConstant.Dot, constant.AnyField),
	)
}
