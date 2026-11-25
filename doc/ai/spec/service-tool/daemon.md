---
base: pkg/tool/goatlassiand
---

# Daemon

## The daemon

Every daemon has one `daemon.Daemon`, small ones included. It is the only
owner `Run` holds: it builds everything else and closes what it built. It
does no work of its own.

The daemon takes configuration - the option, a logger, a reporter, and a metric
registry when its workers export gauges - and builds the rest: the database
(`relational.Open`, `migrate.AutoMigrate`, then each store over it,
`database.md`), its clients, pools and peer guards, then each service from the
parts that service uses. It holds each thing it owns in one field - the same
guard in two fields reads as an owner never closed. Its `Close` closes
everything it built and nothing it was given. A daemon that owns nothing - a
reader over the local system, paths and a logger - has no `Close`, and `Run`
defers nothing.

A daemon is never handed a finished client, in a parameter or an option
function: whoever closes a handed client closes something it did not
build. A client enters as a factory on the option - a function returning
the client's face, set in `option.New()` to the real constructor, or set
in `Main` when the environment decides whether there is one. The daemon
calls it once and closes what it returns when that has a `Close`. A client
the environment may leave out - a factory returning nil, a configuration
that is off - the daemon builds only when present; the services that use
it then answer `not_configured`, and a surface asks the service rather
than testing a field for nil. A pooled client's factory takes the
instance - `func(*instance.Instance) (face.<Name>Source, error)` - the
build function `pool.New` takes. An option holds what `Main` decides once:
addresses, tokens, a parsed inventory, and factories, which own nothing
until the daemon calls them. Nothing live sits on an option.

A daemon with required peers has a `Wait` that waits on each required
peer's guard in turn. `Run` calls it right after building the daemon and
panics on its error, before the lifecycle starts, so the daemon serves
only once every required peer answers; a peer only a clock reads is
guarded and not waited on (`error-handling/infrastructure.md`, Required
and guarded peers).

A face describes one client in full and its mock implements all of it: the
upstream client's face is what the factory returns, the daemon client's is
`face.Client` (`generated-api.md`). Code shared between daemons and tools
takes a role interface - only the methods it calls, beside the full face
(`FloatSource`), filled by the client, its mock or a service's guarded
method through a function type in a sub-package of the face
(`float_source_function.FloatSourceFunction`). No other
face is partial or mirrors another; neither the daemon nor a service has
one.

## Services

The use cases live in services. A service holds only what the daemon built
it with - the stores, clients, guards and clock it uses, and the other
services it calls - and never the daemon: the daemon holds the service, so
the service holding the daemon is an import cycle. A service calling
another is built with it as a constructor parameter, so the daemon builds
them in dependency order and the calls between services run one way:

```go
// daemon/new.go
m := member.New(s.member)
b := book.New(s.catalog, o.Random)
l := loan.New(s.catalog, m, b, d.clock)
```

State only one service uses - a waiter map, a cache it reads, who is
watching - is that service's own field. A view gathering several services -
a member's page with their loans and holds - is a service of its own, built
with the services it reads, named for what it shows.

The service validates. A surface parses its request into plain values, calls one
operation, and answers the error's class (`error-handling/README.md`); a check
that lived in a route moves into the service, where every surface shares it. An
update gets the record, applies the change, checks and saves - a change applied
inside the store hides the result from the check - and sets only the fields it
was sent. A write that may still be refused is staged and moved into place once
every check passed. A value joined into a path, a URL, a selector or a script is
checked as the narrowest shape it can be - a bare name, a bare number - since an
encoded `%2e%2e` passes `http.ServeMux`'s cleaning. A surface holds the services
it calls and nothing beside them - no client, store or guard of its own to reach
around them, and never the daemon. The daemon keeps no method that only forwards
to a service: each surface is handed the service itself. The daemon's work
belongs in a service, never in a worker or the daemon: a tick or an event calls
one operation that returns its outcomes.

A service reads its paths and credentials through the option, never from the
home directory or the keychain itself, so a test can never reach the real one.

A service test builds that service from its parts and the services it was
built with; a surface test goes served through `Mount` (`testing/double.md`).
Nothing builds a surface over hand-wired services.

## Stores

A store borrows the mapper the daemon opened and has no `Close`
(`database.md`, Store Shape). It saves and fires nothing: the service
notifies after a write that succeeded. One store serves a daemon whose tables
commit independently. Where some tables must commit together and others
never do, the store splits along those consistency boundaries - one store
per aggregate, `store/<aggregate>/` - and a transaction never spans two.
Tables one service owns outright get its own store; tables several services
write in one transaction stay one store, shared by every service writing
them.

A sequence that must commit together runs through the store's
`Atomically`. The store opens, commits and rolls back; the service names
the steps and passes each the transaction's store:

```go
// store/catalog/atomically.go
func (s *Store) Atomically(apply func(*Store) error) error {
    return s.mapper.Transaction(
        func(t *gorm.DB) error {
            return apply(New(t))
        },
    )
}

// service/loan/return.go
func (s *Service) Return(identifier int) error {
    return s.catalog.Atomically(
        func(t *catalog.Store) error {
            if e := s.close(t, identifier, s.clock()); e != nil {
                return e
            }

            return s.restock(t, identifier)
        },
    )
}
```

A step inside reads the store it was passed, never the service's own
field. No service imports the mapper's package or opens a transaction
itself, and nothing calls a peer between `Atomically`'s open and commit:
what a sequence reads from a peer it reads before, and what it tells a
peer it tells after the commit.

## Workers

Workers implement `face.Worker` (`Start()` + `Close()`) and close only
what they started or built - a service or client they were given
belongs to whoever built it. See `lifecycle.md` for details.

A worker that holds state or a connection - a peer guard, a runner a
service launches work into, a socket - is built by the daemon and reaches
back only through function values it was given at construction, never the
daemon: the daemon holds the worker, so the worker holding the daemon is an
import cycle.

A worker that holds nothing but time is a clock. The daemon builds it over
one service operation per tick - goclauded's
`worker.New(v.Reconcile, l, r, constant.ReconcileInterval)` on
`pkg/worker` - and it keeps only what is its own: the ticker, its gauges,
its stop channel and the recovery around the tick. Its `Close` cancels the
context the tick was handed and waits for a tick in flight, so nothing the
daemon closes next is still in use; what a cancelled tick returns is not
reported. No service holds a clock.

The daemon hands its workers out through `Workers()`, guards first, so
they run before anything calls through them, then its clocks:

```go
// daemon/workers.go
func (d *Daemon) Workers() []face.Worker {
    return []face.Worker{
        d.catalogGuard,
        worker.New(d.loan.Remind, d.logger, d.reporter, constant.RemindInterval),
    }
}
```

The lifecycle closes them, and the daemon closes what it built again, so
their `Close` is safe twice (`lifecycle.md`).

## Daemon and CLI Split

A daemon `go<tool>d` pairs with a CLI `go<tool>`, which talks to it
through the daemon client in `<toold>/client/` - the one client other
daemons use too (`generated-api.md`, The Daemon Client). A rule deciding
what the daemon's data means - which threads count as unresolved - is a
daemon operation, never the CLI's, so every caller gets the same answer.
