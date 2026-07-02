# Multi-Instance Services

Pattern for services that connect to multiple backends of the same
type (multiple Proxmox hypervisors, multiple Prometheus servers,
multiple PostgreSQL databases). The service manages a registry of
named instances and routes requests to the correct backend.

## When to Use

When a daemon bridges a third-party service and there may be more
than one deployment of that service to connect to. The inventory
allows a single daemon deployment to serve all instances.

Services using this pattern: goprometheusd, goalertmanagerd,
gopostgresd, goproxmoxd.

## Inventory

YAML file with named instances. Each instance carries its own
connection details and credentials. Loaded at startup via
`--inventory` flag with a default of `go<tool>d.yaml`.

```
pkg/tool/go<tool>d/inventory/
├── instance.go         # Instance struct with yaml tags
├── inventory.go        # Inventory struct: Instances []Instance
└── load.go             # Load(path string) *Inventory
```

The Instance struct fields vary by service but always include
`Name` and `Host`. Credentials belong in the instance - not in
environment variables - so each instance authenticates independently.

## Service Layer

The service layer sits between the surfaces (MCP, REST, CLI) and
the upstream clients. It owns three concerns: lazy client creation,
session-scoped instance selection, and instance resolution.

```
pkg/tool/go<tool>d/service/
├── service.go              # struct: inventory, clients map, sessions sync.Map, mu sync.Mutex
├── new.go                  # New(*Inventory) *Service
├── client.go               # Client(instance string) - lazy creation with mutex
├── instance.go             # Instance(name string) - lookup by name
├── instances.go            # Instances() - full list
├── resolve_instance.go     # ResolveInstance(explicit string) - defaulting logic
├── active_instance.go      # ActiveInstance(sessionID string) - MCP session state
└── set_active_instance.go  # SetActiveInstance(sessionID, instance string)
```

### Instance Resolution

`ResolveInstance(explicit string) (string, error)` encodes the
defaulting rule:

- If `explicit` is non-empty, validate it exists and return it
- If exactly one instance is configured, return it
- Otherwise, return an error ("no instance selected" or "multiple
  instances configured, selection required")

All three surfaces call this instead of doing their own resolution.
The rule is: when there's no ambiguity, don't require selection.

### Lazy Client Creation

`Client(instance string) (*<upstream>.Client, error)` creates
upstream clients on first use, caching them in a mutex-protected
map. The client is built from the instance's connection details
and credential fields.

## MCP Surface

### Conditional Tool Registration

When exactly one instance is configured, `list_instances` and
`use_instance` are not registered - there's nothing to list or
choose. When multiple instances exist, both tools appear.

```go
func (s *Server) register() {
    if len(s.service.Instances()) > 1 {
        s.server.AddTool(/* list_instances */)
        s.server.AddTool(/* use_instance */)
    }
    // register all other tools
}
```

### Instance Resolution in Handlers

Every handler resolves the instance before getting a client:

```go
instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

if e != nil {
    return response.Fail("%s", e)
}

c, e := s.service.Client(instance)
```

`activeInstanceName(x)` returns the session's selected instance or
empty string if none was selected. With one instance, the empty
string resolves to the default. With multiple instances and no
selection, `ResolveInstance` returns an error.

### Server Instructions

The identity's `WithInstructions()` is set at package init before
the instance count is known. Use a generic instruction that works
for both cases:

```
"Proxmox hypervisor bridge. Call list_instances to see available instances."
```

When there's one instance the tool won't exist, but the instruction
is still accurate - there's just nothing to list.

## REST Surface

The `instance` query parameter is optional in the OpenAPI spec.
When omitted, `ResolveInstance("")` applies the same defaulting
rule. When multiple instances exist, omitting it returns a 400
`Error` - no Sentry, because it's a caller mistake, not
an infrastructure failure. See `error-handling/rest.md` for the
tier 1 vs tier 2 distinction on strict server.

## CLI Surface

`--instance` is an optional persistent flag. The client wrapper
stores the value (empty if not provided) and passes it to the
generated client on every call. The REST server resolves it via
`ResolveInstance`.

## Deployment

The inventory file is mounted as a Kubernetes file secret:

```yaml
containers:
- name: <tool>
  args: [--inventory, /etc/go<tool>/go<tool>d.yaml]
  volumeMounts:
  - {name: inventory, mountPath: /etc/go<tool>, readOnly: true}
volumes: [{name: inventory, secret: {secretName: <tool>-fsec}}]
```

Credentials live inside the inventory YAML, not in environment
variables. The file secret is the single source for all instance
configuration.

Sentry and telemetry environment variables remain as ConfigMap
and Secret envFrom - they configure the daemon itself, not the
upstream connections.
