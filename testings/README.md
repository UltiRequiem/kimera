# Kimera Test Suite

This directory contains the test suite for Kimera.js. Tests are organized by the permissions they require.

## Directory Structure

Tests are organized into subdirectories based on the permissions they need:

- **`basic/`** - Tests that don't require any special permissions
  - Permission denial tests
  - Basic functionality tests
  - Examples that don't interact with the system

- **`fs/`** - Tests requiring filesystem access (`--fs` flag)
  - File reading tests
  - File writing tests
  - Filesystem permission tests

- **`net/`** - Tests requiring network access (`--net` flag)
  - HTTP fetch tests
  - Network permission tests

- **`env/`** - Tests requiring environment variable access (`--env` flag)
  - Environment variable read/write tests
  - Environment permission tests

- **`combined/`** - Tests requiring multiple permissions (`--fs --env`)
  - Tests that use multiple system capabilities
  - Integration tests

## Running Tests

### Run All Tests

The CI system automatically runs all tests with appropriate permissions:

```bash
go build -o kimera
# Tests are run automatically by directory
```

### Run Tests by Category

Run tests in a specific directory with the appropriate flags:

```bash
# Filesystem tests
./kimera run testings/fs/file-read.js --fs

# Network tests
./kimera run testings/net/fetch-get.js --net

# Environment tests
./kimera run testings/env/permission-env-allowed.js --env

# Combined permission tests
./kimera run testings/combined/permission-demo.js --fs --env

# Basic tests (no flags needed)
./kimera run testings/basic/globals.js
```

## Adding New Tests

When adding new tests, place them in the appropriate directory:

1. **Determine required permissions** - What system access does your test need?
   - Filesystem access? → `fs/`
   - Network access? → `net/`
   - Environment variables? → `env/`
   - Multiple permissions? → `combined/`
   - No permissions? → `basic/`

2. **Create your test file** - Use descriptive names:
   - Good: `file-read-binary.js`, `fetch-headers.js`
   - Avoid: `test1.js`, `mytest.js`

3. **Follow test conventions**:
   - Tests should exit with code 0 on success
   - Use `console.log()` for output
   - Use descriptive messages (e.g., "PASS: Test description")
   - Handle errors appropriately

4. **No manual CI updates needed** - The CI automatically discovers and runs tests based on directory structure!

## Test File Naming

Use descriptive, kebab-case names that indicate what the test does:

- ✓ `file-read.js` - Clear and descriptive
- ✓ `fetch-post-json.js` - Specific about what it tests
- ✓ `permission-fs-denied.js` - Clear about expected behavior
- ✗ `test.js` - Too generic
- ✗ `myNewTest.js` - Not consistent with project style

## Permission Flags

- `--fs` - Enables filesystem read/write operations
- `--net` - Enables network/HTTP operations
- `--env` - Enables environment variable access

Tests are automatically run with the appropriate flags based on their directory.
