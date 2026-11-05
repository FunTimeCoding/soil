# Refactor

How a structural change to a daemon or a shared client is done: survey,
change, prove, gate. The rules the code must meet are in the specs; this
is the procedure.

## Survey

- Read the surfaces - the REST server, the MCP server, the web package -
  not only `run.go`. A surface calling a client directly, and a test
  passing nil, are what a `run.go` read misses. Each conversion so far
  turned up a defect nobody was looking for: a leak, a write blanking
  fields it was not sent, a delete reported that never happened
- A claim put forward as a decision rests on code that was read; say
  what was not read
- Count a method's callers by receiver across every module and build tag
  in one pass: a `go/packages` program resolving each selector through
  `TypesInfo.Selections`, counting calls through the face and inside the
  package too, or a twin only those reach reads as dead. A name-level
  grep only proves a zero. A method gaining an error return still
  compiles where a caller discards the result - errcheck finds those
- Read a peer client's constructor before placing a start-up lookup: a
  constructor that already calls the peer has decided it
- Size a test migration by its doubles, not its files

## Change

- One change at a time, each leaving soil and every consumer building
  and tested. A consumer keeps building through a transitional variadic
  while its call sites move new-beside-old
- A full face over a large client is generated: a program reading the
  client's AST writes the face and trivial mock files as new files,
  refusing to overwrite stateful ones. Must twins copy the client's body;
  zero values cover the basic types (`{}` for a struct, `0` for a float);
  a body returning only the failure carries no import block. Where a
  face exists, only the missing methods go to a scratch file for Edit
- Moving a face or mock home is gosourced work: `move_package` then
  `rename_package` for a mock, `move_symbol` then `rename_symbol` for a
  face
- gosourced refuses every operation while the module does not
  type-check, and loads a consumer together with soil, so keep the
  consumer type-checking between steps. It renames struct methods, not
  interface methods - Edit the face first - and moves a type with
  methods only through `extract_type`. Adding a parameter everywhere is
  `name_parameters` (name the blank `_ context.Context`) then
  `add_parameters`; two calls in order when one parameter feeds another
- Edit changes existing files, one unique anchor per site. In an OpenAPI
  spec the components reference the same schemas as the operations, so
  anchor on the operation's own request schema. A `replace_all` rewrites
  the name inside string literals too. A program writes new files only;
  rewriting YAML through a library re-indents every list

## Prove

- Break by `go test -overlay`, which leaves the tracked tree alone, so a
  proof can run in another session's package while it works there
- A script running breaks compares each target with its backup and
  refuses a break that changed nothing, keeps backups under names unique
  in the backup directory, reads a build failure as no verdict, and reads
  `panic:` and a package `FAIL` as failures - checked by a control break
  that must survive. A break that blocks hangs to the package timeout and
  names no test, so a test of a blocking operation fails by name after a
  bound
- A coverage test's overlay undoes one rename consistently in face,
  client, mock and callers, so only the test names the method. A face
  test's break removes a mock method no other mock file calls. A round
  trip turns a renamed method into a build failure first
- A mock with one failure for every method fails a sequence at its first
  step; fail one method by its method value (`c.Fail(c.DeviceTypes, e)`),
  which a rename follows

## Gate

The build of soil and every consumer against the local checkout; the
affected suites under `-race`, every consumer's when a shared client's
lifecycle or face changed; `go vet` under every build tag; and the lint
stages run from source (`lint.md`). Staging while a break stands stages
it, so name the files.
