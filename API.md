# Kimera.js API Documentation

This document provides comprehensive documentation of all APIs available in
Kimera.js.

## Command Line Interface

### `kimera`

Launch the interactive REPL (Read-Eval-Print Loop).

```sh
kimera
```

**Behavior:**

- Starts an interactive JavaScript console
- Displays version information on startup
- Evaluates JavaScript expressions in real-time
- Maintains a single QuickJS runtime for the entire session (efficient)
- Exit with `Ctrl+C`, `Ctrl+D`, or by calling `close()`

**Example:**

```sh
$ kimera
Kimera 0.2.0
Exit using Ctrl+C or Ctrl+D

> const x = 10
undefined
> x * 2
20
> console.log("Hello!")
Hello!
undefined
> const add = (a, b) => a + b
undefined
> add(5, 3)
8
```

### `kimera run [file]`

Execute a JavaScript or TypeScript file.

```sh
kimera run <file> [flags]
```

**Arguments:**

- `file` (required) - Path to the JavaScript (.js) or TypeScript (.ts) file to
  execute

**Flags:**

- `--fs` - Allow filesystem access (required for `Kimera.readFile()` and `Kimera.writeFile()`)
- `--net` - Allow network access (required for `fetch()`)
- `--env` - Allow environment variable access (required for `Kimera.getEnv()` and `Kimera.setEnv()`)

**Examples:**

```sh
# Run a JavaScript file (no permissions)
kimera run script.js

# Run a TypeScript file with type annotations
kimera run app.ts

# Run with permission flags
kimera run script.js --fs              # Filesystem only
kimera run script.js --net             # Network only
kimera run script.js --env             # Environment variables only
kimera run script.js --fs --net --env  # All permissions
```

**Error Handling:**

The command returns proper exit codes:

- `0` - Success
- `1` - Error (file not found, syntax error, runtime error)

Error messages are formatted with context:

```sh
$ kimera run missing.js
Error: failed to run file: failed to read file "missing.js": open missing.js: no such file or directory
```

### `kimera version`

Display the version of Kimera.js.

```sh
kimera version
```

**Output:**

```
Kimera 0.2.0
```

### `kimera help` / `kimera --help`

Display help information about available commands.

```sh
kimera --help
```

**Output:**

```
Kimera is a modern JavaScript/TypeScript runtime that provides a REPL
and file execution capabilities.

Usage:
  kimera [flags]
  kimera [command]

Available Commands:
  help        Help about any command
  run         Run a JavaScript or TypeScript file
  version     Print the version information

Flags:
  -h, --help   help for kimera

Use "kimera [command] --help" for more information about a command.
```

## JavaScript/TypeScript APIs

### Console API

#### `console.log(...args)`

Prints messages to standard output.

**Parameters:**

- `...args` - Any number of arguments to print

**Example:**

```javascript
console.log("Hello");
console.log("Multiple", "arguments");
console.log("Value:", 42);
console.log(`Template: ${1 + 1}`);
```

### Kimera File System API

The `Kimera` object provides file system operations.

#### `Kimera.readFile(filePath)`

Reads the entire contents of a file as a string.

**Parameters:**

- `filePath` (string) - Path to the file to read

**Returns:**

- (string) - The file contents as a string

**Throws:**

- Error if file doesn't exist or cannot be read

**Example:**

```javascript
// Read a text file
const content = Kimera.readFile("./data.txt");
console.log(content);

// Read with error handling
try {
  const data = Kimera.readFile("./config.json");
  console.log(data);
} catch (error) {
  console.log("Failed to read file");
}
```

#### `Kimera.writeFile(filePath, content)`

Writes content to a file, creating it if it doesn't exist or overwriting if it
does.

**Parameters:**

- `filePath` (string) - Path to the file to write
- `content` (string) - Content to write to the file

**Returns:**

- null

**Throws:**

- Error if file cannot be written

**Example:**

```javascript
// Write simple text
Kimera.writeFile("output.txt", "Hello World!");

// Write multi-line content
const data = "Line 1\nLine 2\nLine 3";
Kimera.writeFile("multiline.txt", data);

// Overwrite existing file
Kimera.writeFile("existing.txt", "New content");

// Write with error handling
try {
  Kimera.writeFile("/protected/file.txt", "data");
} catch (error) {
  console.log("Permission denied");
}
```

**Note:** File system operations require the `--fs` flag:

```sh
kimera run script.js --fs
```

### Kimera Environment Variables API

The `Kimera` object provides environment variable operations.

#### `Kimera.getEnv(varName)`

Reads the value of an environment variable.

**Parameters:**

- `varName` (string) - Name of the environment variable to read

**Returns:**

- (string) - The value of the environment variable, or an empty string if not set

**Throws:**

- Error if environment variable access is not allowed (requires `--env` flag)

**Example:**

```javascript
// Read an environment variable
const path = Kimera.getEnv("PATH");
console.log("PATH:", path);

// Check if a variable exists
const myVar = Kimera.getEnv("MY_VAR");
if (myVar) {
  console.log("MY_VAR is set to:", myVar);
} else {
  console.log("MY_VAR is not set");
}

// With error handling
try {
  const home = Kimera.getEnv("HOME");
  console.log("Home directory:", home);
} catch (error) {
  console.log("Environment access denied");
}
```

#### `Kimera.setEnv(varName, value)`

Sets the value of an environment variable in the current process.

**Parameters:**

- `varName` (string) - Name of the environment variable to set
- `value` (string) - Value to set the environment variable to

**Returns:**

- null

**Throws:**

- Error if environment variable access is not allowed (requires `--env` flag)
- Error if the operation fails

**Example:**

```javascript
// Set an environment variable
Kimera.setEnv("MY_VAR", "my_value");

// Verify it was set
const value = Kimera.getEnv("MY_VAR");
console.log(value); // "my_value"

// Set with error handling
try {
  Kimera.setEnv("CONFIG_PATH", "/etc/config");
  console.log("Environment variable set successfully");
} catch (error) {
  console.log("Failed to set environment variable");
}
```

**Note:** Environment variable operations require the `--env` flag:

```sh
kimera run script.js --env
```

### Global Functions

#### `close()`

Exits the Kimera runtime immediately.

**Example:**

```javascript
console.log("Before exit");
close();
console.log("This won't print");
```

### Global Objects

Kimera provides the following global objects:

- `console` - Console logging API
- `Kimera` - File system, environment variables, and runtime API
- `close()` - Function to exit the runtime

## Language Support

### JavaScript

Kimera.js supports modern JavaScript (ES6+) including:

- **Arrow Functions**
  ```javascript
  const add = (a, b) => a + b;
  ```

- **Template Literals**
  ```javascript
  const name = "World";
  console.log(`Hello, ${name}!`);
  ```

- **Async/Await**
  ```javascript
  const asyncFunc = async () => {
    return "result";
  };

  asyncFunc().then(console.log);
  ```

- **Destructuring**
  ```javascript
  const obj = { a: 1, b: 2 };
  const { a, b } = obj;
  ```

- **Spread Operator**
  ```javascript
  const arr = [1, 2, 3];
  const arr2 = [...arr, 4, 5];
  ```

- **For...of Loops**
  ```javascript
  for (const item of [1, 2, 3]) {
    console.log(item);
  }
  ```

### TypeScript

Full TypeScript support with automatic transpilation:

```typescript
// Type annotations
const greet = (name: string): string => {
  return `Hello, ${name}!`;
};

// Interfaces
interface User {
  name: string;
  age: number;
}

const user: User = {
  name: "Alice",
  age: 30,
};

// Generics
function identity<T>(arg: T): T {
  return arg;
}
```

**Note:** TypeScript is transpiled to JavaScript using esbuild, so type checking
happens at transpilation time, not runtime.

## Error Handling

All file operations and runtime errors can be caught using try-catch:

```javascript
try {
  const content = Kimera.readFile("missing.txt");
} catch (error) {
  console.log("Error occurred:", error);
}

try {
  Kimera.writeFile("/root/protected.txt", "data");
} catch (error) {
  console.log("Permission denied");
}
```

## Best Practices

1. **Always handle file errors**: Use try-catch blocks when reading or writing
   files
   ```javascript
   try {
     const data = Kimera.readFile(filename);
   } catch (error) {
     console.log("File operation failed");
   }
   ```

2. **Use async/await for cleaner code**: Kimera supports modern async patterns
   ```javascript
   const main = async () => {
     // Your async code here
   };

   main();
   ```

3. **Leverage TypeScript for type safety**: Use .ts files for better development
   experience
   ```typescript
   interface Config {
     port: number;
     host: string;
   }
   ```

## Architecture & Implementation Details

### Error Handling

Kimera uses Go's idiomatic error handling patterns:

- All errors are properly wrapped with context using `fmt.Errorf` with `%w`
- Error chains can be inspected using `errors.Is` and `errors.As`
- Clear error messages indicate the operation that failed and why

Example error flow:

```
Error: failed to run file: failed to read file "script.js": open script.js: no such file or directory
       [cmd layer]          [core layer]           [os layer]
```

### Resource Management

- QuickJS runtimes and contexts are properly freed using `defer`
- REPL maintains a single runtime for the entire session (performance
  optimization)
- Each `kimera run` command creates a fresh runtime (isolation)

### TypeScript Transpilation

- TypeScript files are transpiled to JavaScript using esbuild
- Transpilation happens at runtime before execution
- No separate build step required
- Type checking occurs during transpilation

## Limitations

Current limitations of Kimera.js:

- **No module system**: import/export statements not yet supported
- **No npm packages**: Cannot install or use npm packages
- **Limited Node.js compatibility**: Only a subset of APIs available
- **Single-threaded**: No worker threads or parallel execution
- **No DOM APIs**: This is not a browser runtime
- **HTTP/Fetch API**: Currently experimental with limited functionality

## Examples

### Reading and Processing a File

```javascript
// Read a file and count lines
const content = Kimera.readFile("data.txt");
const lines = content.split("\n");
console.log(`File has ${lines.length} lines`);
```

**Run with:**

```sh
kimera run script.js --fs
```

### Creating a Log File

```javascript
const logMessage = (message) => {
  const timestamp = new Date().toISOString();
  const logEntry = `[${timestamp}] ${message}\n`;

  try {
    // Append to log (read existing, append, write back)
    let existingLog = "";
    try {
      existingLog = Kimera.readFile("app.log");
    } catch (e) {
      // File doesn't exist yet, that's okay
    }
    Kimera.writeFile("app.log", existingLog + logEntry);
  } catch (error) {
    console.log("Failed to write log");
  }
};

logMessage("Application started");
logMessage("Processing data");
```

**Run with:**

```sh
kimera run script.js --fs
```

### TypeScript with File Operations

```typescript
interface FileData {
  content: string;
  lines: number;
}

const analyzeFile = (path: string): FileData => {
  const content: string = Kimera.readFile(path);
  const lines: number = content.split("\n").length;

  return { content, lines };
};

const result = analyzeFile("data.txt");
console.log(`Lines: ${result.lines}`);
```

**Run with:**

```sh
kimera run script.ts --fs
```
