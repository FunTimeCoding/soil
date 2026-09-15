# File Structure

- **One identity per file** - each function, method, or type lives in
  its own file, named as the snake_case of the identity:
  `addFileToTar()` → `add_file_to_tar.go`, `type Store struct` →
  `store.go`. The `file_identity` analyzer enforces it; gosourced
  `extract_to_file` is the instrument for splitting.
- **Test files don't split** - one `_test.go` per feature, living in
  the domain root's `unit/` with the subpackage prefix (see
  `../test-placement.md`).
- **One struct with receivers per package** - see
  `../package-design.md` for the full rule and extraction pattern.
