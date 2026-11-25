# Outage Tests

A guarded daemon with a REST or MCP surface has an outage table
(`go:pkg/web/outage_tester`): every REST operation and MCP tool against a peer
that is down, each on a fresh daemon built by the daemon's
`<path>/unit/outage_server` over one shared reporter, so no operation inherits
a guard another one opened. The double fails with
`go:pkg/web/outage_tester.RefusedDial`. The table declares the daemon's usual
reach and the names that differ from it (`go:pkg/web/constant.OutageReach`): a
reaching name opens the outage and
answers `unavailable` - a 503, or a tool error carrying the guard's own outage
text - a degrading one opens it and answers anyway, one elsewhere leaves the
guard closed, and none captures. An exception or override naming nothing
served fails by name.

## Traps

- **The guard is found by name.** `serve` hands back the daemon's own guard's
  `Unavailable` through `Tester.Guard`, which fails when no guard carries the
  peer's name. An index into `Workers()` hands a table expecting elsewhere a
  healthy guard that passes whatever the names do.
- **An empty fixture never reaches the peer.** An operation stopping at
  `not_found` never opens the guard and fails as never reaching it - override
  a placeholder argument the operation refuses, or seed in `serve` the state it
  needs when one write without the peer serves it. State the peer must create,
  or a fixture that grows per tool, is no table's to build: the service test
  reaches that path with the peer stopped, and the table declares the name for
  what the placeholder meets.
- **An unstarted runner waits forever.** A provisioning runner never started
  waits on `sync`, so `serve` closes it before the table calls anything.

## What proves the guard

A quiet class reads the same guarded or not, so an outage test cannot prove
the guard; what only the guard does can - the next call failing at once,
without asking the peer. The double fails as the client does: a dial failure
is a `connection.Failure` (a hand-built `net.OpError` classifies only around
a real errno, `os.NewSyscallError("connect", syscall.ECONNREFUSED)`), where a
bare `unreachable` error passes through as a refusal and the guard never
opens. A test pointing a real client at a down peer, with no double between,
gives it `go:pkg/web/outage_tester.Tester.RefusedAddress`, a closed server's
address. The guard answers `unavailable` on the call that opens the outage and
reports it only after `Lasting`, so a fast test counts the peer's calls, not
the reporter's events. A guard comes back only through its probe loop, on a
real timer from `ProbeMinimum`, so a test that needs the peer back fails a
call with a fault the guard passes through untouched.
