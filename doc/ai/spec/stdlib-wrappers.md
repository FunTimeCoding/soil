# Stdlib Wrappers

Commonly used stdlib functions have PanicOnError wrappers in shared
packages. Use the wrapper instead of the raw call + manual error
handling. This keeps call sites clean and routes all error behavior
through `errors.PanicOnError`.

## Filesystem - `pkg/system/`

| Wrapper | Wraps | Notes |
|---------|-------|-------|
| `system.Open(path)` | `os.Open` | |
| `system.Create(path)` | `os.Create` | |
| `system.MakeDirectory(path)` | `os.MkdirAll` | |
| `system.ReadDirectory(path)` | `os.ReadDir` | Returns `[]os.DirEntry` |
| `system.ReadBytes(base, name)` | `fs.ReadFile` via embed | Takes base + name, not full path |
| `system.ReadFile(base, name)` | Same, returns `string` | |
| `system.ReadLine(r)` | `r.ReadString('\n')` | |
| `system.ReadAll(reader)` | `io.ReadAll` | Generic readers; HTTP response bodies go through `web.ReadBytes` (panics typed `unreadable_body`) |
| `system.Stat(path)` | `os.Stat` | Returns `os.FileInfo` |
| `system.FileStat(f)` | `f.Stat()` on `*os.File` | |
| `system.RelativePath(base, target)` | `filepath.Rel` | |
| `system.Copy(dst, src)` | `io.Copy` | |
| `system.TarWriteHeader(...)` | `tar.FileInfoHeader(...)` | |
| `system.Remove(path)` | `os.RemoveAll` | Recursive - use for directories |
| `system.RemoveFile(path)` | `os.Remove` | Single file only |
| `system.Move(from, to)` | `os.Rename` | |

**Touch pattern in tests:** `errors.PanicClose(system.Create(path))`
creates an empty file and closes it in one line.

Other stdlib idioms with a panicking home: `errors.PanicFlush(w)` for
`w.Flush()`, `defer errors.PanicClose(f)` for `defer f.Close()`,
`_, e = f.Write(b); errors.PanicOnError(e)` for `f.Write(b)`, and
`result, _, e := transform.String(t, s); errors.PanicOnError(e)` for
`transform.String`.

When the raw stdlib call is used for flow control (checking existence
with `os.IsNotExist`, scanning directories that may not exist yet),
the wrapper is not appropriate - the caller needs the error.

## JSON - `pkg/notation/`

| Wrapper | Wraps | Notes |
|---------|-------|-------|
| `notation.Marshal(v)` | `json.Marshal` | Returns `[]byte` |
| `notation.Decode(s, &v)` | `json.Unmarshal` via string | Returns `error` |
| `notation.MustDecode(s, &v, verbose)` | Same + PanicOnError | For trusted input (DB columns, config) |
| `notation.DecodeBytes(b, &v)` | `json.Unmarshal` | Returns `error` |
| `notation.MustDecodeBytes(b, &v, verbose)` | Same + PanicOnError | For trusted input |

Use `MustDecode` / `MustDecodeBytes` for data from your own
systems (database columns, API responses from your own services).
Use the non-Must variants when the error outcome changes control
flow (try-parse patterns, validating external input).

For HTTP response bodies: `notation.MustDecodeBytes(system.ReadAll(resp.Body), &v, false)`.

## Time - `pkg/time/`

| Wrapper | Wraps | Notes |
|---------|-------|-------|
| `time.Parse(layout, s)` | `time.Parse` | Returns `time.Time` |

Import as `"github.com/funtimecoding/soil/pkg/time"` - shadows
stdlib `time` when the stdlib package is no longer needed directly.

## Web, MCP, and process patterns

The same doctrine above stdlib: prefer the soil helper over the raw
call plus boilerplate.

| Raw pattern | Soil replacement |
|---|---|
| `w.Header().Set(...); json.NewEncoder(w).Encode(v)` | `web.EncodeNotation(w, v)` |
| `json.NewEncoder(w).Encode(v)` (header set separately) | `web.Encode(w, v)` |
| `w.Header().Set(constant.ContentType, constant.Object)` | `web.ObjectHeader(w)` |
| `m.HandleFunc("GET /x", h)` / `fmt.Sprintf("GET %s", path)` | `m.HandleFunc(route.Get(path...), h)` (`pkg/web/route`, variadic parts, `route.Post` for POST) |
| `mcp.NewToolResultError(fmt.Sprintf(...))` | `response.Fail(...)` |
| `mcp.NewToolResultText(notation.MarshalIndent(v))` | `response.SuccessAny(v)` |
| `mcp.NewToolResultText("message")` | `response.Success("message")` |
| `exec.Command(name, args...)` | `run.New().Start(name, args...)` |
| `t.Format(...)` for human-facing display | `time.FormatCompact(t)` (`pkg/time`) - today renders clock-only, older renders date and minute |
| `html.Td(...FormatCompact(t)...)` in dashboard tables | `layout.TimeCell(t)` (`pkg/web/layout`) - compact, small, nowrap; date-only cells take `layout.TimeCellClass` |

## When Not to Wrap

- `os.Stat` for existence checks (`os.IsNotExist`) - flow control
- `os.ReadDir` on directories that may not exist - flow control
- `os.Remove` in batch cleanup where failure is self-healing - use
  `hub.CaptureException` instead of panic
- `json.Unmarshal` in try-parse patterns - flow control
- Generated code (`client/generated.go`) - don't modify
