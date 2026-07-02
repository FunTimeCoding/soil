# Coding Conventions

## Error Handling

- **Prefer `PanicOnError` over returning errors** - see `error-handling/` for the full strategy and the exceptions (MCP handlers, flow control)
- **Prefer `PanicClose` over checking close errors**
- Use `LogClose` only in loops/uncertain contexts (e.g., validation)
- **Error variable naming:**
  - `e` for first error in scope
  - `f` for second error in same scope (avoid shadowing)
  - `g` for third error in same scope
  - When running out of e/f/g in deeply nested scopes: `createFail`, `saveFail` (prefer short words like "fail" over "error")
  - Eliminate variables entirely when possible: `errors.PanicOnError(someCall())`
  - Descriptive names like `createFail` or `printFail` are still "smells" - indicates code needs refactoring

## Testing Philosophy

- **Unix approach: silence on success, noise on failure**
- No progress logs (`t.Log("doing X...")`)
- No success logs (`t.Logf("✓ X succeeded")`)
- Only `t.Fatalf()` for must-succeed setup, `t.Errorf()` for validation failures
- **Assert package** - use `pkg/assert` for all test assertions. Argument order is `(t, expected, actual)` - expected first, actual second. Type-specific functions: `assert.String`, `assert.Integer`, `assert.Float`, `assert.True`, `assert.NotNil`, `assert.StringContains`, etc. Never use testify.
- Extract helper functions, make them panic instead of returning errors

## Code Style

- **Never-nesting ideal** - avoid deep indentation
- **No defensive programming in tests** - "should work or blow up"
- Remove explanatory comments if code is self-documenting
- Helper functions have minimal/no error returns in test code
- **Short variable names:**
  - `e`/`f`/`g` for errors in nested scopes
  - `h` for headers, `i` for FileInfo in loops
  - `z` for gzip writers, `t` for tar writers, `f` for files
  - `path` over `outputPath`, `name` over `fileName`
  - Single letters for obvious types in local variables and struct receivers (the type already documents the receiver). Function parameters prefer descriptive names - callers read the signature without seeing the body.
  - Avoid `-ing` and plural forms in names
  - Prefer shorter words: "fail" over "error", "path" over "filePath"
- **String concatenation** - use `join.Empty(a, b)` (`pkg/strings/join`) for pure string-string joins. Use `fmt.Sprintf` only when there's actual formatting (numbers, padding, mixed types). For long prose strings (descriptions, help text), use a single long line rather than multiline `"foo" + "bar"` literal concatenation.
- **Function/method chaining** - pass return values directly: `p.WritePKGINFO(p.CreateDataTar())`
- **Blank identifier for byte counts** - ignore bytes read/written when not needed: `content, _ := io.ReadAll(tr)`, `_, f = outFile.Write(data)`
  - Never ignore error returns, only byte counts
- **`new(expr)` for pointers** (Go 1.26) - use `new(time.Now())` instead of `t := time.Now(); &t`. Works for any expression: `new(true)`, `new("value")`, `new(42)`
- **Defer placement** - immediately after resource creation and error check:
  ```go
  file, e := os.Open(path)
  errors.PanicOnError(e)
  defer errors.PanicClose(file)
  ```
- **Empty line placement:**
  - Blank lines go BEFORE and AFTER entire control structure blocks (around the braces)
  - **Return statement blank lines:**
    - At function level: blank line before return (unless return is the only statement in entire function)
    - In control structure blocks (if/for/switch): blank line before return ONLY if multiple statements in that block
  - Only control structures and return statements break consecutive simple statement chains
  - No blank lines at function/method start or end (first/last statements sit flush with braces)
  - Consecutive simple statements (assignments, function calls, defer) should be compact (no blank lines)
  - Multiline statements (spanning multiple lines) count as one statement for blank line purposes
  - Aim for functions/methods with ~15 calls or fewer to maintain readability
  - Line length: break at 80 characters, prioritizing fewer total lines (break single-arg calls before multi-arg calls)
- **Multiline function calls** - when args don't fit one line, each arg gets its own line with closing paren separate
- **Prefer go-library wrappers** - use go-library wrappers instead of stdlib + error checking. Quick reference:

  | stdlib call | go-library replacement |
  |---|---|
  | `os.Open(path)` | `system.Open(path)` |
  | `os.Create(path)` | `system.Create(path)` |
  | `os.MkdirAll(path, 0755)` | `system.MakeDirectory(path)` |
  | `filepath.Rel(base, target)` | `system.RelativePath(base, target)` |
  | `r.ReadString('\n')` | `system.ReadLine(r)` |
  | `os.Stat(path)` | `system.Stat(path)` |
| `f.Stat()` (on `*os.File`) | `system.FileStat(f)` |
  | `io.Copy(dst, src)` | `system.Copy(dst, src)` |
  | `tar.FileInfoHeader(...)` | `system.TarWriteHeader(...)` |
  | `w.Flush()` | `errors.PanicFlush(w)` |
  | `defer f.Close()` | `defer errors.PanicClose(f)` |
  | `f.Write(b)` | `_, e = f.Write(b); errors.PanicOnError(e)` |
  | `transform.String(t, s)` | `result, _, e := transform.String(t, s); errors.PanicOnError(e)` |
  | `w.Header().Set(...); json.NewEncoder(w).Encode(v)` | `web.EncodeNotation(w, v)` |
  | `json.NewEncoder(w).Encode(v)` (header set separately) | `web.Encode(w, v)` |
  | `w.Header().Set(constant.ContentType, constant.Object)` | `web.ObjectHeader(w)` |
  | `mcp.NewToolResultError(fmt.Sprintf(...))` | `response.Fail(...)` |
  | `mcp.NewToolResultText(notation.MarshalIndent(v))` | `response.SuccessAny(v)` |
  | `mcp.NewToolResultText("message")` | `response.Success("message")` |
  | `exec.Command(name, args...)` | `run.New().Start(name, args...)` |

  **Touch pattern in tests:** `errors.PanicClose(system.Create(path))` creates an empty file and closes it in one line.

## Structure

- **One-file-per-function** - each function or method lives in its own file. File name is the snake_case form of the function name: `addFileToTar()` → `add_file_to_tar.go`, `httpFail()` → `http_fail.go`. Test files don't split - one `_test.go` per feature, named after the primary source file.
- **One-type-per-file** - struct definitions go in a file named after the struct (snake_case): `type Store struct` → `store.go`. One struct with receivers per package - see `package-design.md` for the full rule and extraction pattern.

### Constants

- **Constants in `constant.go`** - exported constants get their own `constant.go` file in the package. Prefer a single flat `constant.go` per package. When an iota enum is defined with a named type, both the type and the const values belong in the same constant package - separating them creates circular imports.
- **Reuse existing constants** - use `web/constant.Listen`, `argument.Name`, `parameter.Query`, `parameter.Limit`, etc. Never hardcode strings that already have a constant in the codebase. MCP parameter names shared across tools live in `generative/model_context/parameter/`.
- **Constant value naming** - abbreviations and acronyms are avoided in constant values, not just names. `Link = "link"` not `Link = "url"`. The value should be the honest, full-word form.
- **Semantic constant splitting** - when the same string serves different layers (DB column, HTML form field, domain key), use separate constants even if the values are identical today. Example: `NameFieldKey = "name"` (field system), `NameColumn = "name"` (GORM), `NameParameter = "name"` (web routes/forms). They could diverge independently.
- **Propagate generic constants toward go-library** - when a constant is used across multiple tools or packages (e.g. `FormMethod = "method"` for HTML forms), it belongs in go-library's shared vocabulary (e.g. `web/constant`), not duplicated in each consumer.
- **Option struct naming** - named after the domain concept, not `Option` (e.g., `option.Log`, `option.Build`, `option.Commit`). File named after the struct. Constructor in `new.go` returns pointer.
- Tests use `assert.*` helpers to reduce nesting
- **`new_environment.go`** - when a package reads from environment variables, the environment constructor lives in `new_environment.go` and calls `New()` after reading the required vars. Pattern: `func NewEnvironment() *Client { return New(environment.Required(constant.HostEnvironment), environment.Required(constant.TokenEnvironment)) }`. The base `New()` always takes explicit params; `NewEnvironment()` is the env-reading wrapper.
- **`example/` subpackage** - every client package should have an `example/` when there's a real usage scenario to demonstrate. Don't create empty or contrived examples just to fill the slot. Entry point lives in `cmd/example/<name>/main.go`. Inactive examples are kept in `if false {}` blocks rather than deleted.
- **Constructors** - `New()` in `new.go`, zero-arg or with params. One per package. Zero-arg `New()` is appropriate when call sites use different field combinations (callers set fields after construction). `Stub()` in `stub.go` returns a zero-value pointer - used for GORM model references (`AutoMigrate`, `Model`, `Delete`) and error return paths (when a method fails and needs to return something alongside the error). Never use `Stub()` for real construction - that's `New()`. See `package-design.md` for when types warrant their own packages.
- **Local tool installation** - use `gobuild --copy-to-bin <name>` to build and install tools locally. Never use `go install`. Procfiles use `go run`.
- Fail fast with stacktraces (Sentry integration)
- Helpers return only success values, panic on errors
- Prefer extracting logic to testable helpers over inline test code

## Forbidden Imports

Enforced by the `goanalyze` `forbidden_import` analyzer (AST-based, no false positives on string literals).

| Forbidden | Use instead |
|-----------|-------------|
| `"flag"` | `"github.com/spf13/pflag"` |
| `"github.com/stretchr/testify"` | `"github.com/funtimecoding/go-library/pkg/assert"` |

## Import Aliases

- **No acronyms** in package names or import aliases
- Alias only when package names collide in the same file
- Alias by the subsystem name, not the role in the current file:
  - `generative` for `pkg/generative/model_context/server` (when it collides with the tool's own `server/`)
  - `generated` for the oapi-codegen `server/` package (when it collides with another `server/`)
  - `webConstant` for `pkg/web/constant` (when it collides with a local `constant/`)
- **The more local package keeps the natural name** - when two packages share a last segment (e.g., both called `server`), the tool's own package is imported unaliased; the shared infrastructure package gets the alias. In files that only import one of them, no alias is needed.

## Interfaces

- **Generic interfaces** (no domain type imports) - define in `pkg/face/`. These are like `io.Reader`: small, shared, used by many consumers. Examples: `Worker`, `Notifier`, `Clock`.
- **Domain-specific interfaces** (reference types like `*alert.Alert`) - define at the consumer. Domain types often already import `face`, so putting the interface in `face` would create an import cycle. Example: `AlertSource` in `worker/worker.go`.
- **Mock convention** - hand-rolled structs in `mock_*` sub-packages near the real implementation (e.g., `mock_notifier/`, `mock_client/`). One file per method. Methods to manipulate state for testing (e.g., `Add()`, `Remove()`).

## Deviations

**vs Claude defaults:** Claude writes defensive, verbose tests with every error checked and logged

**vs Idiomatic Go:** Idiomatic Go returns `(T, error)` everywhere; we prefer panic in non-public APIs
