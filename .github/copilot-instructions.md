# Kimera.js - GitHub Copilot Instructions

## Project Overview

Kimera.js is a minimal JavaScript/TypeScript runtime written in Go. It is built on top of QuickJS and esbuild, providing a lightweight alternative to Node.js and Deno.

## Tech Stack

- **Language**: Go (1.17+)
- **JavaScript Engine**: QuickJS (via github.com/lithdew/quickjs)
- **TypeScript Support**: esbuild (github.com/evanw/esbuild)
- **CLI Framework**: Cobra (github.com/spf13/cobra)

## Project Structure

- `/cmd` - Command-line interface implementations
- `/core` - Core runtime functionality
  - `/std` - Standard library implementations for JavaScript
- `/testings` - JavaScript/TypeScript test files

## Coding Guidelines

### Go Code

- Use tabs for indentation (Go standard)
- Follow standard Go formatting (use `go fmt`)
- Use meaningful variable names (prefer `errorReadingFile` over generic `err`)
- Always use `defer` for cleanup operations (e.g., `defer ctx.Free()`)
- Lock OS thread when working with QuickJS: `runtime.LockOSThread()`
- Handle errors explicitly - check and report errors appropriately
- Use `cobra.Command` pattern for CLI commands

### JavaScript/TypeScript Tests

- Place test files in the `/testings` directory
- Use `.js` or `.ts` extensions
- Tests should be executable scripts that exit with status 0 on success
- Use descriptive filenames (e.g., `file-read.js`, `async.ts`)

## Building and Testing

### Build

```bash
go build
```

### Run Tests

The project uses a custom test runner that executes JavaScript/TypeScript files in the `/testings` directory:

```bash
./kimera run testings/<test-file>.js
```

### Testing with Go

```bash
go test ./...
```

Note: Currently, there are no Go unit tests. If adding Go tests, place them alongside the code they test with `_test.go` suffix.

## Features to Implement Carefully

### Runtime Capabilities

- TypeScript is automatically transpiled via esbuild
- File system access is opt-in via `--fs` flag
- Network access is opt-in via `--net` flag
- Environment variables access is opt-in via `--env` flag

### Global APIs

The runtime provides custom global APIs through the `Kimera` namespace:
- File operations: `Kimera.readFile()`, `Kimera.writeFile()`
- These are injected via the `__dispatch` function

## Security Considerations

- Keep permissions opt-in (fs, net, env flags)
- Validate user input in CLI commands
- Be careful with QuickJS memory management - always Free() contexts and runtimes
- Don't expose system internals through JavaScript APIs

## CI/CD

- GitHub Actions workflows in `.github/workflows/`
- Build workflow runs on every push/PR
- Test workflow executes all JavaScript/TypeScript files in `/testings`
- Release process managed by GoReleaser

## Dependencies

- Keep dependencies minimal
- Update dependencies carefully as they affect runtime security
- CGO is required (CGO_ENABLED=1) for QuickJS bindings

## Common Patterns

### Adding a New CLI Command

```go
var newCmd = &cobra.Command{
    Use:   "command [args]",
    Short: "Brief description",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        core.HandleCommand(args[0])
    },
}
rootCmd.AddCommand(newCmd)
```

### Adding JavaScript Globals

1. Implement the function in Go
2. Register it via `globals.Set("name", ctx.Function(handler))`
3. Expose it through the standard library in `/core/std/`

## Documentation

- Keep README.md up to date with new features
- Document CLI flags and commands
- Provide examples for new JavaScript APIs
- Use clear, concise language

## Error Handling

- Use `CheckJSError` helper for JavaScript errors
- Exit with appropriate codes (0 for success, 1 for errors)
- Provide helpful error messages to users
- Don't panic unless absolutely necessary

## Future Considerations

- The project is working towards automatic builds (issue #2)
- Keep cross-platform compatibility in mind (Linux, macOS, Windows)
- Consider ARM64 support (already in GoReleaser config)
