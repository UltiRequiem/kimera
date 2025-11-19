// TypeScript test for file handling
const message: string = "Hello from TypeScript!";
Kimera.writeFile("/tmp/typescript-test.txt", message);

const readMessage: string = Kimera.readFile("/tmp/typescript-test.txt");
console.log(readMessage);

if (readMessage === message) {
  console.log("TypeScript test passed!");
}
