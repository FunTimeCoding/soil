---
base: doc/ai/spec
---

# Multi-Instance Services

Pattern for daemons that connect to several backends of the same type -
several Proxmox hypervisors, Prometheus servers or PostgreSQL databases. One
deployment serves every instance named in an inventory and routes each
request to the right backend.

The plumbing lives in `pkg/inventory/` and its subpackages; a daemon supplies
its instance type and a function that builds a client. The daemons on it are
the ones that call `inventory.Load`: `grep -rl 'inventory.Load\[' pkg/tool`.

## When to Use

When a daemon bridges a third-party service and there may be more than one
deployment of that service to reach.

## The Instance Type

The daemon's `instance` package holds the instance struct. Its fields vary by
service, and credentials belong in the instance - not in environment
variables - so each instance authenticates independently.

- `InstanceName()` and `InstanceHost()` on the pointer receiver make it an
  inventory entry. `InstanceHost()` returns what `list_instances` should show,
  wherever the struct keeps it.
- An optional `Validate() error` runs per instance at load.
- `type Inventory = inventory.Inventory[Instance, *Instance]` in the same
  package, so the rest of the daemon writes `*instance.Inventory`.

Rules that span instances - goproxmoxd's unique `Index` per instance - are
daemon-side functions over the inventory, called in `Main` after `Load`.

## The Inventory File

YAML with an `instances` list, read by `inventory.Load` from the `--inventory`
flag, which defaults to `Identity.InventoryPath()`. A missing file, malformed
YAML or a failed `Validate` panics at startup. `inventory.New` builds one in
memory for tests.

In a deployment the file is mounted read-only, typically from a secret, and
`--inventory` points at it. It is the single source of instance
configuration; the environment still carries the daemon's own settings, such
as error reporting and telemetry.

## Resolution

Every surface resolves an instance name through the same rule, in
`inventory.Resolve`:

- An explicit name that is not in the inventory is `not_found`
- No instances configured is `not_configured`
- Exactly one instance is the default
- More than one with no selection is `not_selected`

When there is no ambiguity, selection is not required.

## The Service

The service holds the inventory, a `go:pkg/generative/model_context/selection`
`Selection` and one `pool.Pool` per kind of client, and exposes the inventory
and selection as
`Inventory()` and `Selection()`, plus an accessor for each pool something
outside the service needs. Surfaces use those directly; the service adds no
wrappers for listing, finding or resolving instances.

A pool builds each instance's client on first use from the daemon's build
function and caches it. A failed build is not cached, so the next call
retries; a name outside the inventory is `not_found`. When a pool's clients
hold connections, `Run` defers that pool's `Close()`, which closes every
client the pool built that has a `Close()`. `Set` puts a mock in place for
tests; a client given that way stays its owner's to close.

## MCP Surface

`instance_tool.New(v.Inventory(), v.Selection(), r)` is held by the MCP server.
`Register` adds `list_instances` and `use_instance` only when two or more
instances are configured - with one there is nothing to choose. Handlers
resolve before touching a client and answer a resolution error by its class
through `answer` - `not_found` and `not_selected` fail plainly,
`not_configured` and a context without a session are captured:

```go
instance, e := s.instances.Resolve(x)

if e != nil {
    return s.answer(e)
}
```

### Sessions

`use_instance` stores the choice against the MCP session. The server is built
with `WithSessionEnd(v.Selection().Clear)`, so a session's choice is
forgotten when it ends - on DELETE, when an SSE stream closes, or when the
idle sweeper reaps a client that left without either. A client holding its
event stream open is never reaped; the heartbeat keeps it fresh.

Selection rides on the session of the legacy protocol versions, which soil's
MCP setup pins. Protocol version 2026-07-28 drops protocol-level sessions, so
moving to it needs another home for the choice.

### Server Instructions

The identity carries the instance sentence as a conditional paragraph, so a
single-instance deployment does not tell the model to call tools it does not
have:

```go
WithInstructions(
    "Prometheus query interface.",
    paragraph.New(
        inventoryConstant.MultiInstance,
        inventoryConstant.MultiInstanceInstruction,
    ),
)
```

The MCP server renders it with
`instance_tool.Instructions(constant.Identity, len(v.Inventory().Instances))`.

## REST Surface

`instance` is an optional query parameter. The handler resolves it with
`Inventory().ResolveInstance`, so omitting it follows the same rule. A
resolution error answers by its class, as the table in
`error-handling/README.md` gives it: an instance that names nothing is
`not_found` and answers 404, a missing selection is `not_selected` and
answers 400, and a deployment with no instance at all is `not_configured`,
a captured 500. Every resolving operation declares its 404.

## CLI Surface

`--instance` is an optional persistent flag. The client passes it on every
call, and the REST server resolves it.

## Testing

A two-instance MCP test drives the real service through `inventory.New`,
`model_context_server` and `model_context_client`. It covers the refusal
before `use_instance`, `list_instances` decoded into
`instance_response.Response` with the active flag on the chosen one, and -
where the daemon has a mock client - a tool following the selected instance's
client, with a mock per instance set into the pool. Probe it by dropping
`Register`: the test must fail.
