package report

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
)

func classify(
	p *entry.Entry,
	mapping map[string]string,
) string {
	body := remap(p.From.Body, mapping)
	same := p.From.Key() == p.To.Key()

	switch {
	case same && remap(p.From.Signature, mapping) != p.To.Signature:
		return constant.ClassSignature
	case same && body != p.To.Body:
		return constant.ClassBody
	case same:
		return constant.ClassUnchanged
	case p.From.Member() != p.To.Member():
		return constant.ClassRenamed
	case body == p.To.Body:
		return constant.ClassMoved
	default:
		return constant.ClassMovedChanged
	}
}
