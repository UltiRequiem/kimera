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

- REPL 👇

```sh
kimera
```

- TypeScript out of the box

### Installation

Not yet released, working on automatic builds on
[#2](https://github.com/UltiRequiem/kimera/issues/2).

```bash
go install github.com/UltiRequiem/kimera@latest
```

### License

Kimera.js is Licensed under the MIT License.
