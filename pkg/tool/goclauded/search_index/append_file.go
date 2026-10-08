package search_index

import (
	"bufio"
	"errors"
	library "github.com/funtimecoding/soil/pkg/errors"
	"io"
	"os"
)

func (x *Index) appendFile(
	session string,
	path string,
) {
	f, e := os.Open(path)

	if errors.Is(e, os.ErrNotExist) {
		return
	}

	library.PanicOnError(e)
	defer library.PanicClose(f)
	i, e := f.Stat()
	library.PanicOnError(e)
	consumed, turn := x.position(session)

	if consumed == i.Size() {
		return
	}

	if consumed > i.Size() {
		_, e = x.remove(session)
		library.PanicOnError(e)
		consumed, turn = 0, ""
	}

	_, e = f.Seek(consumed, io.SeekStart)
	library.PanicOnError(e)
	t, e := x.database.Begin()
	library.PanicOnError(e)
	defer rollback(t)
	r := bufio.NewReader(f)

	for {
		line, g := r.ReadBytes('\n')

		if errors.Is(g, io.EOF) {
			break
		}

		library.PanicOnError(g)
		consumed += int64(len(line))
		entries, next := extract(line, session, turn)
		turn = next

		for _, n := range entries {
			insert(t, n)
		}
	}

	_, e = t.Exec(
		`INSERT INTO harbor_file (session, consumed, turn) VALUES (?, ?, ?)
		ON CONFLICT(session) DO UPDATE SET consumed = excluded.consumed, turn = excluded.turn`,
		session,
		consumed,
		turn,
	)
	library.PanicOnError(e)
	library.PanicOnError(t.Commit())
}
