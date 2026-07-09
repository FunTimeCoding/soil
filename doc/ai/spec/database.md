# Database

Convention for service persistence: sqlite by default, postgres by
locator, the choice made in one place and announced in the log. A
service binary always has a working zero-config database; deployments
opt into postgres explicitly.

## Storage Location

Service data lives at the identity-derived path. Never hand-join
`~/.local/share` segments.

- `system.StorageDirectory(name, create)` - `~/.local/share/<name>`;
  pass `create` true only with write intent
- `constant.Identity.StorageDirectory(create)` - the same, from the
  tool's identity
- `constant.Identity.LitePath()` - `~/.local/share/<name>/<name>.sqlite`

Path derivation is pure - directories come into existence where
writing starts: the database openers (`lite.New` and the raw-sql
store constructors) create the parent of their path, so flag
defaults and `--help` never touch the filesystem, and a custom
`--lite` path gets its directory for free.

## Driver Rule

Enforced by the goanalyze `restricted_call` analyzer: `gorm.Open`
and `sql.Open` may only be called in the `relational` packages -
services and tests go through the openers below.

One sqlite driver per access style, both pure Go:

- gorm stores use `glebarez/sqlite` through `relational/lite` - never
  `gorm.io/driver/sqlite`, which wraps the CGO driver and breaks under
  cross-compilation (gobuild builds with CGO disabled)
- raw `database/sql` stores open through
  `relational/lite/connection.New`, which carries the glebarez driver
  registration

## Openers

Six constructors, each returning what its consumer actually uses:

- `lite.New(l, path) *gorm.DB` - sqlite; logs the backend; WAL and
  foreign_keys arrive as DSN `_pragma` parameters, so every pooled
  connection gets them (an `Exec` would only reach one)
- `connection.New(l, path) *sql.DB` - the raw twin of `lite.New` for
  stores that speak `database/sql` directly; logs; same DSN parameters
- `connection.NewMemory() *sql.DB` - in-memory for raw-store tests;
  a named shared-cache database, because raw query-while-iterating
  patterns deadlock on a pinned single connection and a plain
  `:memory:` is one database per pooled connection
- `lite.NewMemory() *gorm.DB` - in-memory sqlite for tests;
  `foreign_keys = ON`, pool pinned to one connection (a pooled second
  connection would open its own empty database), no WAL
- `relational.NewMapper(l, locator) *gorm.DB` - postgres, mapper
  only; logs; the postgres twin of `lite.New`
- `relational.New(locator) *Database` - the full object: pgx pool,
  `database/sql`, mapper. For provisioning tools and services that
  query outside gorm; silent - a service using it as its store logs
  `relational.PostgresMessage` itself (gonixd is the one case)

A service that only speaks gorm receives only a mapper. Switching it
to the full `Database` later is a one-line change at its edge.

## Selection

`relational.Open(l, locator, litePath) *gorm.DB` encodes the rule:
postgres when the locator is set, else sqlite at the lite path, panic
when neither. The openers themselves log the backend (`database:
postgres` or `database: sqlite`) - the witness line is emitted where
the connection is born, never maintained by hand at call sites. The
log is the witness when a deployment forgets its locator.

Wiring in `Main()` - `a.Database()` registers `--lite` (chain:
`LITE_PATH` env > `Identity.LitePath()`) and `--postgres` (chain:
`POSTGRES_LOCATOR` env > empty):

```go
a.Database()
a.Parse(version, gitHash, buildDate)
o := option.New()
o.PostgresLocator = a.GetString(argument.Postgres)
o.LitePath = a.GetString(argument.Lite)
```

And in `Run()`:

```go
l := logger.New(context.Background())
s := store.New(relational.Open(l, o.PostgresLocator, o.LitePath))
```

Every gorm-backed service is dual-backend through this wiring -
deployments choose postgres by setting the locator in their
environment Secret, everything else lands on the lite default and says so in
the log. Raw-sql sqlite stores (gomemoryd, goqueryd) register only
`a.Lite()` and emit the same log line via `relational.LiteMessage`.
The line states the backend class and nothing more - the path and
locator are environment configuration, not log content. The memory
openers are silent - test noise is not essential logging.

## Store Shape

Backend resolution happens at the edge; the store is backend-agnostic.
The constructor takes the mapper and owns migration:

```go
func New(m *gorm.DB) *Store {
    errors.PanicOnError(m.AutoMigrate(NewClick()))

    return &Store{mapper: m}
}
```

Tests build the same store over a throwaway file:

```go
s := store.New(lite.New(filepath.Join(t.TempDir(), constant.TestDatabase)))
```

## Deployment

`POSTGRES_LOCATOR` is set per service, as a sealed secret mounted via
`envFrom`, pointing at that service's own database. The env name is
generic but the value is service-specific - a DSN with a database name
in it - so it must never be exported globally (shell profiles, .envrc):
any dual-backend binary launched from that shell would connect to the
wrong service's database and run its AutoMigrate there.
