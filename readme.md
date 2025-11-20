# Kimera.js

A minimal JavaScript/TypeScript runtime written in Go.

Built on top of [QuickJS](https://github.com/bellard/quickjs) and
[esbuild](https://github.com/evanw/esbuild), Kimera provides a fast, lightweight
alternative for running JavaScript and TypeScript code.

## Quick Start

```javascript
const asyncSayHello = async (text) => {
  console.log(text);
};

asyncSayHello("Hello World!");
```

```sh
kimera run myScript.js
Hello World!
```

## Installation

```bash
go install github.com/UltiRequiem/kimera@latest
```

## Features

Kimera.js provides a lightweight JavaScript/TypeScript runtime with the
following capabilities:

### 1. Interactive REPL

Launch an interactive Read-Eval-Print Loop for testing JavaScript code:

```sh
kimera
```

The REPL supports:

- Multi-line statements
- Variable persistence across commands
- Full JavaScript ES6+ syntax
- Exit using `Ctrl+C` or `Ctrl+D`

### 2. Script Execution

Run JavaScript or TypeScript files directly:

```sh
kimera run myScript.js
kimera run myScript.ts
```

**Permission Flags**:

Kimera implements a secure-by-default permission system. Scripts must explicitly request access to sensitive operations:

```sh
kimera run script.js --fs    # Allow filesystem access
kimera run script.js --net   # Allow network access
kimera run script.js --env   # Allow environment variable access

# Multiple flags can be combined
kimera run script.js --fs --net --env
```

### 3. TypeScript Support

Native TypeScript support out of the box - no configuration needed. Kimera
automatically transpiles TypeScript files using esbuild:

```typescript
const greet = (name: string): void => {
  console.log(`Hello, ${name}!`);
};

greet("World");
```

### 4. Modern JavaScript Features

Full support for modern JavaScript syntax including:

- Async/await
- Arrow functions
- Template literals
- Destructuring
- ES6+ features
- Promises

Example:

```javascript
const fetchData = async () => {
  return "Data loaded";
};

fetchData().then((data) => console.log(data));
```

### 5. Console API

Standard console logging functionality:

```javascript
console.log("Simple message");
console.log("Multiple", "arguments", "supported");
console.log(`Template literals: ${1 + 1}`);
```

### 6. File System API

Built-in file handling through the `Kimera` global object:

#### Reading Files

```javascript
// Read file content as string
const content = Kimera.readFile("path/to/file.txt");
console.log(content);
```

#### Writing Files

```javascript
// Write string content to file
Kimera.writeFile("path/to/output.txt", "Hello World!");

// Write multi-line content
const data = "Line 1\nLine 2\nLine 3";
Kimera.writeFile("output.txt", data);
```

#### Error Handling

File operations throw errors for invalid operations:

```javascript
try {
  const content = Kimera.readFile("nonexistent.txt");
} catch (error) {
  console.log("File not found!");
}
```

### 7. Environment Variables API

Access and modify environment variables through the `Kimera` global object (requires `--env` flag):

#### Reading Environment Variables

```javascript
// Get an environment variable
const path = Kimera.getEnv("PATH");
console.log("PATH:", path);

// Check if a variable exists
const myVar = Kimera.getEnv("MY_VAR");
if (myVar) {
  console.log("MY_VAR is set to:", myVar);
} else {
  console.log("MY_VAR is not set");
}
```

#### Setting Environment Variables

```javascript
// Set an environment variable
Kimera.setEnv("MY_VAR", "my_value");

// Verify it was set
const value = Kimera.getEnv("MY_VAR");
console.log(value); // "my_value"
```

#### Error Handling

Environment operations throw errors when permission is not granted:

```javascript
try {
  const value = Kimera.getEnv("PATH");
} catch (error) {
  console.log("Environment access denied!");
}
```

**Note:** Environment variable operations require the `--env` flag:

```sh
kimera run script.js --env
```

### 8. Global Objects

Kimera provides the following global objects:

- `console` - Console logging API
- `Kimera` - File system, environment variables, and runtime API
- `close()` - Function to exit the runtime

### 9. Version Information

Check the installed Kimera version:

```sh
kimera version
```

### 10. HTTP/Fetch API (Experimental)

Make HTTP requests with the fetch API (requires `--net` flag):

```javascript
// Make a GET request
const response = fetch("https://api.example.com/data");
console.log("Status: " + response.status);
const data = response.json();

// Make a POST request
const postResponse = fetch("https://api.example.com/data", {
  method: "POST",
  body: JSON.stringify({ key: "value" }),
  headers: {
    "Content-Type": "application/json",
  },
});
```

**Note:** Network operations require the `--net` flag:

```sh
kimera run script.js --net
```

## CLI Usage

```sh
# Run the REPL (all permissions enabled)
kimera

# Run a JavaScript file
kimera run script.js

# Run a TypeScript file
kimera run script.ts

# Run with specific permissions
kimera run script.js --fs           # Filesystem access
kimera run script.js --net          # Network access
kimera run script.js --env          # Environment variables
kimera run script.js --fs --net     # Multiple permissions

# Check version
kimera version

# Get help
kimera --help
```

## Development

### Building from Source

```bash
git clone https://github.com/UltiRequiem/kimera.git
cd kimera
go build
```

### Running Tests

```bash
go test ./...
```

## Roadmap

- [ ] Module system (import/export)
- [ ] npm package support
- [ ] More Node.js API compatibility
- [ ] Better error messages
- [ ] Debugger support
- [x] Permission system implementation (--fs, --net, --env flags)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Kimera.js is licensed under the MIT License.
