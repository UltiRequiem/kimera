console.log(this === globalThis);

for (const key in globalThis) {
  console.log(key);
}
