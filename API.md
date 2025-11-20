# Kimera.js API Documentation

This document provides comprehensive documentation of all APIs available in Kimera.js.

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
- Exit with `ctrl+c` or by calling `close()`

**Example:**
```sh
$ kimera
Kimera 0.1.0
exit using ctrl+c or close()
> const x = 10
undefined
> x * 2
20
> console.log("Hello!")
Hello!
undefined
```

### `kimera run [file]`

Execute a JavaScript or TypeScript file.

```sh
kimera run <file>
```

**Arguments:**
- `file` - Path to the JavaScript (.js) or TypeScript (.ts) file to execute

**Flags:**
- `--fs` - Allow file system access (reserved for future use)
- `--net` - Allow network access (reserved for future use)
- `--env` - Allow environment variable access (reserved for future use)

**Example:**
```sh
kimera run script.js
kimera run app.ts
```

### `kimera version`

Display the version of Kimera.js.

```sh
kimera version
```

**Output:**
```
Kimera 0.1.0
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

Writes content to a file, creating it if it doesn't exist or overwriting if it does.

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
- `Kimera` - File system and runtime API
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
  age: 30
};

// Generics
function identity<T>(arg: T): T {
  return arg;
}
```

**Note:** TypeScript is transpiled to JavaScript using esbuild, so type checking happens at transpilation time, not runtime.

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

1. **Always handle file errors**: Use try-catch blocks when reading or writing files
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

3. **Leverage TypeScript for type safety**: Use .ts files for better development experience
   ```typescript
   interface Config {
     port: number;
     host: string;
   }
   ```

## Limitations

Current limitations of Kimera.js:

- No built-in network/HTTP APIs
- No environment variable access (planned)
- No module system (import/export) yet
- Limited Node.js API compatibility
- No npm package support
- Single-threaded execution
- No DOM APIs (this is not a browser runtime)

## Examples

### Reading and Processing a File

```javascript
// Read a file and count lines
const content = Kimera.readFile("data.txt");
const lines = content.split("\n");
console.log(`File has ${lines.length} lines`);
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
