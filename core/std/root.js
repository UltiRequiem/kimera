const console = {
  log(...args) {
    globalThis.__dispatch("console", ...args);
  },
};

function close() {
  globalThis.__dispatch("close");
}

globalThis = {
  ...globalThis,
  console,
  close,
};

globalThis.window = globalThis;
