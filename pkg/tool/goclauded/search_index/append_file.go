package search_index

import (
	"bufio"
	"errors"
	soilErrors "github.com/funtimecoding/soil/pkg/errors"
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

	soilErrors.PanicOnError(e)
	defer soilErrors.PanicClose(f)
	i, e := f.Stat()
	soilErrors.PanicOnError(e)
	consumed, turn := x.position(session)

	if consumed == i.Size() {
		return
	}

	if consumed > i.Size() {
		_, e = x.remove(session)
		soilErrors.PanicOnError(e)
		consumed, turn = 0, ""
	}

	_, e = f.Seek(consumed, io.SeekStart)
	soilErrors.PanicOnError(e)
	t, e := x.database.Begin()
	soilErrors.PanicOnError(e)
	defer rollback(t)
	r := bufio.NewReader(f)

	for {
		line, g := r.ReadBytes('\n')

		if errors.Is(g, io.EOF) {
			break
		}

		soilErrors.PanicOnError(g)
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
	soilErrors.PanicOnError(e)
	soilErrors.PanicOnError(t.Commit())
}
