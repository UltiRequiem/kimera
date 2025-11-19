// Comprehensive test for file handling
console.log("=== Kimera File Handling Test Suite ===");

// Test 1: Basic write
console.log("Test 1: Basic file write");
Kimera.writeFile("/tmp/test1.txt", "Test content 1");
console.log("✓ Write successful");

// Test 2: Basic read
console.log("Test 2: Basic file read");
const content1 = Kimera.readFile("/tmp/test1.txt");
if (content1 === "Test content 1") {
  console.log("✓ Read successful");
} else {
  console.log("✗ Read failed");
}

// Test 3: Overwrite existing file
console.log("Test 3: Overwrite file");
Kimera.writeFile("/tmp/test1.txt", "New content");
const content2 = Kimera.readFile("/tmp/test1.txt");
if (content2 === "New content") {
  console.log("✓ Overwrite successful");
} else {
  console.log("✗ Overwrite failed");
}

// Test 4: Multi-line content
console.log("Test 4: Multi-line content");
const multiLine = "Line 1\nLine 2\nLine 3";
Kimera.writeFile("/tmp/multiline.txt", multiLine);
const readMultiLine = Kimera.readFile("/tmp/multiline.txt");
if (readMultiLine === multiLine) {
  console.log("✓ Multi-line handling successful");
} else {
  console.log("✗ Multi-line handling failed");
}

// Test 5: Error handling for non-existent file
console.log("Test 5: Error handling");
try {
  Kimera.readFile("/tmp/nonexistent-file-12345.txt");
  console.log("✗ Should have thrown error");
} catch (e) {
  console.log("✓ Error handling successful");
}

console.log("=== All tests completed! ===");
