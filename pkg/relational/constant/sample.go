package constant

const (
	FixtureOldTable       = "old_notes"
	FixtureNewTable       = "notes"
	FixtureOldColumn      = "body"
	FixtureNewColumn      = "content"
	FixtureStaleColumn    = "draft"
	FixtureCreateNotes    = "CREATE TABLE old_notes (identifier integer PRIMARY KEY, body text NOT NULL DEFAULT '', draft text)"
	FixtureCreateBoth     = "CREATE TABLE old_notes (identifier integer PRIMARY KEY, body text, content text)"
	FixtureNoteRow        = "INSERT INTO old_notes (identifier, body, draft) VALUES (1, 'kept', 'gone')"
	FixtureRenamedContent = "SELECT content FROM old_notes WHERE identifier = 1"
	FixtureMovedBody      = "SELECT body FROM notes WHERE identifier = 1"
)
