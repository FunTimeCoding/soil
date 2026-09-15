# Statement Style

- **Never-nesting ideal** - avoid deep indentation; guard clauses and
  early returns over nested blocks.
- **Function chaining** - pass return values directly:
  `p.WritePKGINFO(p.CreateDataTar())`.
- **Blank identifier for byte counts only** - ignore bytes
  read/written when not needed: `content, _ := io.ReadAll(tr)`,
  `_, f = outFile.Write(data)`. Never ignore error returns.
- **`new(expr)` for pointers** (Go 1.26) - `new(time.Now())` instead
  of `t := time.Now(); &t`. Works for any expression: `new(true)`,
  `new("value")`, `new(42)`.
- **Defer placement** - immediately after resource creation and error
  check:

  ```go
  file, e := os.Open(path)
  errors.PanicOnError(e)
  defer errors.PanicClose(file)
  ```

- **String concatenation** - use `join.Empty(a, b)`
  (`pkg/strings/join`) for pure string-string joins. Use
  `fmt.Sprintf` only when there is actual formatting (numbers,
  padding, mixed types). For long prose strings (descriptions, help
  text), use a single long line rather than multiline `"foo" + "bar"`
  literal concatenation.
