---
base: doc/ai/spec
---

# Entrypoint Spec

Shared conventions for all `cmd/` programs - `main`, `Main()`, reporter
integration.

## main

Every `cmd/<name>/main.go` only delegates to `Main()`:

```go
package main

import "github.com/funtimecoding/soil/pkg/tool/go<tool>"

func main() {
    go<tool>.Main()
}
```

Nothing is linked in. The version stamp is Go's own build information, read by
`stamp.New()` - see `build.md`.

## Main Function

`Main()` is the first real function in every program. There are two
shapes depending on whether the tool uses flat flags or subcommands.

### Choosing between flat flags and subcommands

**Flat flags** (`argument.NewInstance` + identity): daemons and
single-purpose tools. Daemons should always use flat flags - a
config path, maybe `--port`, nothing more. Whether a daemon takes
`--port` at all depends on its deployment target - see
`service-tool.md`. Standalone tools that do one thing also use
flat flags.

**Subcommands** (cobra + identity): tools with multiple distinct
operations. Typically CLI clients that talk to a daemon (gopostgres,
gonetbox, gohabitica), but also standalone tools that grew past a
single verb.

### Flat-flag Main (daemons, single-purpose tools)

```go
func Main() {
    r := reporter.New(constant.Identity.Name()).Start()
    defer func() { r.RecoverFlush(recover()) }()
    a := argument.NewInstance(constant.Identity)
    a.String(argument.File, "", "File to wait for")
    a.String(argument.Process, "", "Process to wait for")
    a.String(argument.Locator, "", "Locator to wait for")
    a.String(argument.Contains, "", "String for locator")
    a.Duration(argument.Timeout, 3*time.Minute, "")
    a.Boolean(argument.Verbose, false, "Verbose output")
    a.Parse()
    o := option.New()
    o.File = a.GetString(argument.File)
    o.Process = a.GetString(argument.Process)
    o.Locator = a.GetString(argument.Locator)
    o.Contains = a.GetString(argument.Contains)
    o.Timeout = a.GetDuration(argument.Timeout)
    o.Verbose = a.GetBoolean(argument.Verbose)
    wait.Run(o)
}
```

`argument.NewInstance` takes the tool's identity, creates a scoped
flag set, and wires `--help` from the identity. `Parse()` registers
`--version`, parses, and exits cleanly on `--help` or `--version`. It also
registers `--notation` unless the tool already has one; `--version --notation`
prints the stamp as JSON (`pkg/stamp/report`), the same shape the lifecycle
server's version route serves.
After parsing, populate the option struct using instance getters
(`a.GetString`, `a.GetBoolean`, `a.GetDuration`) or required
variants (`a.Required`, `a.RequiredInteger`).

Registration methods: `Boolean`, `String`, `Integer`, `Duration`,
`BooleanVariable`, `StringVariable`, `IntegerVariable`,
`StringSliceVariable`, `BooleanShort`, `IntegerShort`.

Retrieval methods: `GetBoolean`, `GetString`, `GetInteger`,
`GetDuration`, `Required`, `RequiredInteger`, `RequiredPositional`,
`Positionals`, `Argument`, `ArgumentCount`, `Slice`,
`PositionalFallback`.

### Subcommand Main (multi-operation tools)

```go
func Main() {
    s := instrument.NewCommandLine(constant.Identity)
    defer func() { s.Flush(recover()) }()
    c := client.NewEnvironment()
    o := &cobra.Command{
        Use:   constant.Identity.Usage(),
        Short: constant.Identity.Description(),
    }
    o.AddCommand(listItems(c))
    o.AddCommand(createItem(c))
    argument.CobraInstrument(o, s)
    argument.CobraStamp(o, constant.Identity)
    errors.PanicOnError(o.Execute())
}
```

#### Telemetry and the version stamp

`argument.CobraInstrument` records CLI telemetry for every command under
its full path (`gonix character list`): success after the command runs,
error when a panic ends it, through the deferred `Flush`. A root's own
`PersistentPreRun`/`PostRun` keep running. `instrument.NewCommandLine`
waits at most half a second at exit for the record to send, so an
unreachable telemetry host never holds a shell. Commands do not record
individually; an `os.Exit` inside a command skips both the record and the
flush.

`argument.CobraStamp` gives Cobra the build's version: the same
block every tool prints on `--version`, or the JSON report with `--notation`.
Cobra handles `--help` natively. Identity provides `Use` and `Short` on the
root command.
No option struct - each subcommand owns its own flags.

#### One file per subcommand

Each subcommand lives in its own file and returns a `*cobra.Command`:

```go
func createItem(c *client.Client) *cobra.Command {
    var name string
    result := &cobra.Command{
        Use:   "create-item [value]",
        Short: "Create a new item",
        Args:  cobra.ExactArgs(1),
        Run: func(
            _ *cobra.Command,
            arguments []string,
        ) {
            fmt.Println(c.Create(arguments[0], name))
        },
    }
    result.Flags().StringVar(
        &name,
        "name",
        "",
        "Item name",
    )

    return result
}
```

Subcommand flags use `result.Flags().StringVar` (bound to local
variables), not the argument instance. Required flags use
`result.MarkFlagRequired`.

## Mandatory and optional dependencies

A client package reading its own environment offers one or both
constructors:

- `NewEnvironment()` for a dependency the program cannot run
  without. It uses `environment.Required`, so a missing variable
  panics at startup - the correct outcome when the program would be
  useless anyway.
- `NewOptional()` for a dependency the program works without. It
  returns nil when the dependency is not configured, and the caller
  gates on that:

```go
if claude := connector.NewOptional(); claude != nil {
    // construct and register the feature
}
```

The nil propagates as the feature's own absence rather than a
separate flag - a nil connector means no worker, no store, and the
tools it serves never register. One judgment, made where the
knowledge lives, instead of every consumer testing a variable
itself.

Each package names its own discriminator: the variable that carries
no sensible default. For `connector` that is the token, since host
and port fall back to localhost.

Offer `NewOptional` from a package whose consumers are not all
known. A program that controls its own deployment can require
whatever it likes.

## Instrument Integration

Programs that carry telemetry - daemons and subcommand CLIs - create
an instrument at the top of `Main()`, before any other work.
`pkg/instrument` bundles the observability pair - the Sentry reporter
and the telemetry recorder - behind one constructor and one exit
defer.

```go
s := instrument.New(constant.Identity)
defer func() { s.Flush(recover()) }()
```

The Sentry release is the tag the build is on or follows
(`stamp.New().Tag()`), read by the reporter itself.

- `Recorder()` and `Reporter()` expose the halves as `face.Recorder`
  and `face.Reporter` for downstream components
- `RecordCommand(name)` records a successful CLI command; it takes a
  string so the package stays cobra-free
- The recorder reads `GOTELEMETRY_HOST` and `GOTELEMETRY_PORT`
  (falling back to localhost and the 8080 listen default), speaking
  https unless `GOTELEMETRY_INSECURE` is set
- Daemon `Run()` accepts `face.Instrument` and pulls the halves where
  it wires them; the concrete `*instrument.Instrument` never crosses
  `Run()`

## Reporter Integration

Programs without telemetry create the reporter directly. It captures
unhandled panics and provides error reporting to all downstream
components.

```go
r := reporter.New(constant.Identity.Name(), version).Start()
defer func() { r.RecoverFlush(recover()) }()

// ... flag parsing, option construction
Run(o, r)
```

- Reporter is always created - empty locator produces noop behavior
  internally (no branching, no nil checks)
- Reporter is the first thing in `Main()` so the defer runs last
- `RecoverFlush` captures the panic, flushes to Sentry if configured,
  prints the panic value, and exits
- The reporter is passed as a separate `Run()` parameter, not on the
  option struct (option structs hold configuration, not constructed
  dependencies)
- `Run()` accepts `face.Reporter` (the interface), not the concrete
  `*reporter.Reporter`

See `pillars.md` for the full wiring pattern including logger and
recovery middleware.

### Direct exits and the reporter

Some tools call `os.Exit(1)` for expected failure conditions (no results,
validation failures, upload errors). This intentionally bypasses the reporter
defer - these are not crashes and should not be reported as errors. The reporter
covers unexpected panics only.

### Exit codes

Exit codes are `0` for success and `1` for every failure. Nothing
distinguishes failure kinds by number - the response body already
carries that, and `error-handling/rest.md` defines how: a tier 1
rejection returns the `Error` schema with no `event_identifier`,
while tier 2 and 3 return `ErrorResponse` carrying one.

A CLI client that talks to a daemon prints the response body and
exits `1` when the status is 400 or above. Both cases bypass the
reporter: a tier 1 rejection is the caller's mistake and carries no
Sentry capture by design, and a tier 2 or 3 failure was already
captured daemon-side - the `event_identifier` in the body is the
handle for it, so capturing again would duplicate one failure as
two events. Only transport failures reach the reporter, through the
normal panic path.

A Cobra CLI emits and exits through a `pkg/terminal` `Terminal` built in
`Main` from the instrument (`t := terminal.New(s)`) and handed to each
subcommand beside its client. `t.Emit(response)` prints the body to stdout,
or at 400 and above to stderr and exits `1`; `t.Exitln` and `t.Exitf` write
to stderr and exit `1`, `t.Exit` exits with a code. A typed client's
response is checked before its `JSON200` is read - `t.Reject(r.Status(),
r.Body)` when it is nil. `t.Blockln` is for a protocol exit that is neither
success nor failure, such as a hook refusing a tool call: the command records
`blocked`. A token comes from `t.Required`, read once the command runs:
`web.DeferredBearerEditor(t.Required, name)` reads it on the first request,
so a missing token ends the command as `error` and `--help` needs none.
Results go to stdout, failures to stderr, so a pipe receives only results.
Every exit records the running command's outcome in telemetry and flushes it
first - still without the reporter. A command that calls `os.Exit` directly
records nothing; that is the shape the terminal replaces.
