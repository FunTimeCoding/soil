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

const (
	FixtureCreateEvent      = "CREATE TABLE event (name text, happened_at datetime, note text)"
	FixtureInsertEvent      = "INSERT INTO event (name, happened_at) VALUES (?, ?)"
	FixtureInsertNotedEvent = "INSERT INTO event (name, happened_at, note) VALUES (?, ?, ?)"
	FixtureEventNewestFirst = "SELECT name FROM event ORDER BY happened_at DESC"
	FixtureEventAfter       = "SELECT name FROM event WHERE happened_at > ? ORDER BY name"
	FixtureEventStoredTime  = "SELECT CAST(happened_at AS TEXT) FROM event ORDER BY name"
	FixtureEventNote        = "SELECT note FROM event WHERE note IS NOT NULL"
	FixtureEventTime        = "SELECT happened_at FROM event WHERE name = ?"
)
