# Stdlib Wrappers

Commonly used stdlib functions have PanicOnError wrappers in shared
packages. Use the wrapper instead of the raw call + manual error
handling. This keeps call sites clean and routes all error behavior
through `errors.PanicOnError`.

## Filesystem - `pkg/system/`

| Wrapper | Wraps | Notes |
|---------|-------|-------|
| `system.ReadDirectory(path)` | `os.ReadDir` | Returns `[]os.DirEntry` |
| `system.ReadBytes(base, name)` | `fs.ReadFile` via embed | Takes base + name, not full path |
| `system.ReadFile(base, name)` | Same, returns `string` | |
| `system.ReadAll(reader)` | `io.ReadAll` | For HTTP response bodies |
| `system.Stat(path)` | `os.Stat` | Returns `os.FileInfo` |
| `system.Remove(path)` | `os.RemoveAll` | Recursive - use for directories |
| `system.RemoveFile(path)` | `os.Remove` | Single file only |
| `system.Move(from, to)` | `os.Rename` | |

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

## When Not to Wrap

- `os.Stat` for existence checks (`os.IsNotExist`) - flow control
- `os.ReadDir` on directories that may not exist - flow control
- `os.Remove` in batch cleanup where failure is self-healing - use
  `hub.CaptureException` instead of panic
- `json.Unmarshal` in try-parse patterns - flow control
- Generated code (`client/generated.go`) - don't modify
