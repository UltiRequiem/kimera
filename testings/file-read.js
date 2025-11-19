// Test reading an existing file
const content = Kimera.readFile("readme.md");
console.log("Successfully read readme.md");
console.log("First 50 characters:", content.substring(0, 50));
