// Test writing to a file
const testContent = "Hello from Kimera!\nThis is a test file.";
Kimera.writeFile("/tmp/kimera-test.txt", testContent);
console.log("File written successfully!");

// Read it back to verify
const readContent = Kimera.readFile("/tmp/kimera-test.txt");
console.log("File content:", readContent);

// Verify content matches
if (readContent === testContent) {
  console.log("✓ File write and read test passed!");
} else {
  console.log("✗ Content mismatch!");
}
