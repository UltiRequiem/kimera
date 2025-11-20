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

**Permission Flags** (planned for future use):

```sh
kimera run script.js --fs    # Allow filesystem access
kimera run script.js --net   # Allow network access
kimera run script.js --env   # Allow environment variable access
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

### 7. Global Objects

Kimera provides the following global objects:

- `console` - Console logging API
- `Kimera` - File system and runtime API
- `close()` - Function to exit the runtime

### 8. Version Information

Check the installed Kimera version:

```sh
kimera version
```

### 9. HTTP/Fetch API (Experimental)

Make HTTP requests with the fetch API:

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

## CLI Usage

```sh
# Run the REPL
kimera

# Run a JavaScript file
kimera run script.js

# Run a TypeScript file
kimera run script.ts

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
- [ ] Permission system implementation (--fs, --net, --env flags)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Kimera.js is licensed under the MIT License.
