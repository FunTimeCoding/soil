# Entrypoint Spec

Shared conventions for all `cmd/` programs - linker variables, `Main()`, reporter integration.

## Linker Variables

Every `cmd/<name>/main.go` declares linker variables and delegates to `Main()`:

```go
package main

import "github.com/funtimecoding/soil/pkg/tool/go<tool>"

var (
    Version   string
    GitHash   string
    BuildDate string
)

func main() {
    go<tool>.Main(Version, GitHash, BuildDate)
}
```

These are injected by `gobuild` at compile time - see `build.md`.

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
func Main(
    version string,
    gitHash string,
    buildDate string,
) {
    r := reporter.New(constant.Identity.Name(), version).Start()
    defer func() { r.RecoverFlush(recover()) }()
    a := argument.NewInstance(constant.Identity)
    a.String(argument.File, "", "File to wait for")
    a.String(argument.Process, "", "Process to wait for")
    a.String(argument.Locator, "", "Locator to wait for")
    a.String(argument.Contains, "", "String for locator")
    a.Duration(argument.Timeout, 3*time.Minute, "")
    a.Boolean(argument.Verbose, false, "Verbose output")
    a.Parse(version, gitHash, buildDate)
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
`--version`, parses, and exits cleanly on `--help` or `--version`.
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
func Main(
    version string,
    gitHash string,
    buildDate string,
) {
    r := reporter.New(constant.Identity.Name(), version).Start()
    defer func() { r.RecoverFlush(recover()) }()
    c := client.NewEnvironment()
    o := &cobra.Command{
        Use:     constant.Identity.Usage(),
        Short:   constant.Identity.Description(),
        Version: argument.CobraVersion(version, gitHash, buildDate),
    }
    o.AddCommand(listItems(c))
    o.AddCommand(createItem(c))
    errors.PanicOnError(o.Execute())
}
```

Cobra handles `--version` via the `Version` field and `--help`
natively. Identity provides `Use` and `Short` on the root command.
No option struct - each subcommand owns its own flags.

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

## Reporter Integration

Every program creates a reporter at the top of `Main()`, before any other
work. The reporter captures unhandled panics and provides error reporting
to all downstream components.

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

### `os.Exit` and reporter

Some tools call `os.Exit(1)` for expected failure conditions (no results, validation failures, upload errors). This intentionally bypasses the reporter defer - these are not crashes and should not be reported as errors. The reporter covers unexpected panics only.
