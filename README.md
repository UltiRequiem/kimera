# Kimera.js

[![GitMoji](https://img.shields.io/badge/Gitmoji-%F0%9F%8E%A8%20-FFDD67.svg)](https://gitmoji.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
![Lines Of Code](https://img.shields.io/tokei/lines/github.com/UltiRequiem/kimera?color=blue&label=Total%20Lines)
![CodeQL](https://github.com/UltiRequiem/kimera/workflows/CodeQL/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/kimera)](https://goreportcard.com/report/github.com/UltiRequiem/chigo)

A super fast and lightweight JavaScript Runtime for Scripts.

## Getting Started

```javascript
// myScript.js
const asyncSayHello = async (text) => {
  console.log(text);
};

(async () => await asyncSayHello("Hello World!"))();
```

```bash
$ kimera myScript.js
Hello World!
```

## REPL

```bash
kimera
```

### Installation

```bash
go install github.com/UltiRequiem/kimera@latest
```

### License

This project is licensed under the [MIT LICENSE](./LICENSE.md).
