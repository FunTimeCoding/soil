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

- **By-reference for owned types** - a named struct type this module
  owns travels as a pointer, and the scope is the whole chain rather
  than the one site a finding names: parameters, and the slices and
  maps that carry it, not only returns. `value_return` enforces the
  return half; the rest is judgment. Converting forces the constructor
  the codebase already wants, since `&pkg.X{}` trips `struct_literal`
  for owned packages. Generated types stay by value - oapi-codegen
  types cannot be made to travel by pointer. A type that must travel by
  value - a map key, a value compared on a hot path where a pointer
  would allocate each time - is declared with `//goanalyze:by-value`
  and its reason directly above it; a bare directive does not count.
- **String concatenation** - use `join.Empty(a, b)`
  (`pkg/strings/join`) for pure string-string joins. Use
  `fmt.Sprintf` only when there is actual formatting (numbers,
  padding, mixed types). For long prose strings (descriptions, help
  text), use a single long line rather than multiline `"foo" + "bar"`
  literal concatenation.
