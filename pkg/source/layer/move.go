package layer

import "github.com/funtimecoding/soil/pkg/errors"

func (l *Layer) Move(
	from string,
	to string,
) {
	content, e := l.Read(from)
	errors.PanicOnError(e)
	l.Write(to, content)
	l.Remove(from)
	l.moves[from] = to
}
