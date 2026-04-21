# AGENTS.md - Go Bookmark Project

## Project Overview

This is a CLI bookmark manager written in Go. Single module `soda/go-bookmark` targeting Go 1.26.2.

## Build / Lint / Test Commands

### Build
```bash
go build -o bookmark ./main.go
```

### Run
```bash
go run main.go
```

### Run Single Test
```bash
go test ./... -run TestName -v
```

### Run All Tests
```bash
go test ./...
```

### Test Coverage
```bash
go test ./... -cover
```

### Vet
```bash
go vet ./...
```

### Format
```bash
go fmt ./...
```

### Tidy Dependencies
```bash
go mod tidy
```

## Code Style Guidelines

### Formatting
- Use `gofmt` or `goimports` for automatic formatting
- Indent with tabs, not spaces
- Max line length: 120 characters (soft guideline)

### Naming Conventions
- **Packages**: lowercase, short, no underscores (`bookmark`, not `book_mark`)
- **Files**: lowercase with underscores only for tests (`main.go`, `handler_test.go`)
- **Types**: PascalCase (`BookmarkService`, not `bookmark_service`)
- **Functions exported**: PascalCase; unexported: camelCase
- **Variables**: camelCase; short-lived: short names (`b`, `url`)
- **Constants**: PascalCase for enumerated; camelCase for private
- **Interfaces**: PascalCase, often with `-er` suffix (`Reader`, `Writer`)

### Imports
- Group by: stdlib, then third-party, then internal
- Use `goimports` to manage automatically
- Never import unused packages
- Alias disambiguating packages: `sqlDB "database/sql"`

### Types
- Be explicit with type declarations when helpful for documentation
- Use type aliases for clarity: `type StringMap = map[string]string`
- Pointer receivers for methods that modify the receiver or are large structs
- Value receivers for small read-only methods

### Error Handling
- Return errors as last return value: `func() (T, error)`
- Wrap errors with context: `fmt.Errorf("operation: %w", err)`
- Use sentinel errors for known cases: `var ErrNotFound = errors.New("not found")`
- Never ignore errors with `_`

### Functions
- Keep functions short and focused (< 40 lines target)
- Max 3 parameters; use a struct or options pattern beyond that
- Return early to reduce nesting
- Use named return values sparingly for clarity in longer functions

### Comments
- Comment package-level exports: `// AddBookmark stores a bookmark in the collection.`
- Avoid obvious comments: `// increment i` is noise
- Use辣椒 doc format: `// Package bookmark provides a simple CLI bookmark manager.`

### Concurrency
- Don't leak goroutines; use `sync.WaitGroup`, `context.Context`, or channels
- Share by communication; don't lock shared state unless necessary
- Log fatal concurrency bugs with `runtime.Goexit()` in tests

### Testing
- Test file suffix: `_test.go`
- Test function prefix: `Test` (for `go test`)
- Table-driven tests for multiple cases
- Use `t.Fatal*` / `t.FailNow()` in setup; `t.Error*` for assertion failures

### Project Structure (Simple CLI)
```
go-bookmark/
├── main.go          # entry point
├── go.mod           # module definition
└── *_test.go       # co-located tests (optional)
```

### General Go Idioms
- Zero value initialization is preferred where possible
- Prefer composition over inheritance
- Use interfaces to define contracts (small, focused interfaces)
- Constants via `iota` for enumerated sets
- `go.sum` should be committed; `go.mod` must be committed

## Cursor / Copilot Rules

None present.