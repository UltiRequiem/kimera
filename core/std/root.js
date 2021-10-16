globalThis.console = {
  log: (...args) => {
    globalThis.__dispatch("console", ...args);
  },
};

globalThis.close = () => globalThis.__dispatch("close");

globalThis.window = globalThis;
