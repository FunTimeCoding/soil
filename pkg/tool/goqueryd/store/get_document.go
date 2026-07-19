package store

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"strings"
)

func (s *Store) GetDocument(reference string) (*result.Document, error) {
	var collection, path string

	if strings.HasPrefix(reference, "qmd://") {
		trimmed := strings.TrimPrefix(reference, "qmd://")
		slash := strings.Index(trimmed, separator.Slash)

		if slash == -1 {
			return nil, nil
		}

		collection = trimmed[:slash]
		path = trimmed[slash+1:]
	} else {
		slash := strings.Index(reference, separator.Slash)

		if slash == -1 {
			return nil, nil
		}

		collection = reference[:slash]
		path = reference[slash+1:]
	}

	row := s.database.QueryRow(
		`SELECT d.collection, d.path, d.title, d.hash, c.body
		FROM document d
		JOIN content c ON c.hash = d.hash
		WHERE d.collection = ? AND d.path = ? AND d.active = 1`,
		collection,
		path,
	)
	var d result.Document
	e := row.Scan(&d.Collection, &d.Path, &d.Title, &d.Hash, &d.Body)

	if e == sql.ErrNoRows {
		return nil, nil
	}

	if e != nil {
		return nil, e
	}

	d.VirtualPath = buildVirtualPath(d.Collection, d.Path)
	d.FilePath = join.Empty(d.Collection, separator.Slash, d.Path)
	d.Context = s.ResolveContext(d.Collection, d.Path)

	return &d, nil
}
