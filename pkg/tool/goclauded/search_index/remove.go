package search_index

func (x *Index) remove(session string) (int, error) {
	t, e := x.database.Begin()

	if e != nil {
		return 0, e
	}

	defer rollback(t)
	_, e = t.Exec(
		"DELETE FROM block_text WHERE rowid IN (SELECT rowid FROM block WHERE session = ?)",
		session,
	)

	if e != nil {
		return 0, e
	}

	r, e := t.Exec("DELETE FROM block WHERE session = ?", session)

	if e != nil {
		return 0, e
	}

	count, e := r.RowsAffected()

	if e != nil {
		return 0, e
	}

	_, e = t.Exec("DELETE FROM harbor_file WHERE session = ?", session)

	if e != nil {
		return 0, e
	}

	return int(count), t.Commit()
}
