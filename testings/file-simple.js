// Simple file handling test
console.log("Writing test file...");
Kimera.writeFile("/tmp/test.txt", "Hello World!");

console.log("Reading test file...");
const content = Kimera.readFile("/tmp/test.txt");

console.log(content);
console.log("Test passed!");
