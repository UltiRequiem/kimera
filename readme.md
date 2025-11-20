# Kimera.js

A minimal JavaScript/TypeScript runtime written in Go.

It is built on top of [quickjs](https://github.com/bellard/quickjs) and
[esbuild](https://github.com/evanw/esbuild).

## Getting Started

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

## Features

Kimera.js provides a lightweight JavaScript/TypeScript runtime with the following capabilities:

### 1. Interactive REPL

Launch an interactive Read-Eval-Print Loop for testing JavaScript code:

```sh
kimera
```

The REPL supports:
- Multi-line statements
- Variable persistence across commands
- Full JavaScript ES6+ syntax
- Exit using `ctrl+c` or `close()`

### 2. Script Execution

Run JavaScript or TypeScript files directly:

```sh
kimera run myScript.js
kimera run myScript.ts
```

### 3. TypeScript Support

Native TypeScript support out of the box - no configuration needed. Kimera automatically transpiles TypeScript files using esbuild:

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

fetchData().then(data => console.log(data));
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

### Installation

Not yet released, working on automatic builds on
[#2](https://github.com/UltiRequiem/kimera/issues/2).

```bash
go install github.com/UltiRequiem/kimera@latest
```

### License

Kimera.js is Licensed under the MIT License.
