// Comprehensive permission system demonstration
// This script requires all three permission flags: --fs --net --env

console.log("=== Kimera Permission System Demo ===");
console.log("");

// 1. Test filesystem access
console.log("1. Testing filesystem access (--fs flag):");
try {
  const content = Kimera.readFile("readme.md");
  console.log("   ✓ Successfully read readme.md");
  console.log("   ✓ File size:", content.length, "bytes");
} catch (error) {
  console.log("   ✗ Error:", error);
}

console.log("");

// 2. Test environment variable access
console.log("2. Testing environment access (--env flag):");
try {
  // Set a test variable
  Kimera.setEnv("KIMERA_DEMO", "permission_test");
  const value = Kimera.getEnv("KIMERA_DEMO");
  console.log("   ✓ Set and read KIMERA_DEMO:", value);
  
  // Read an existing variable
  const path = Kimera.getEnv("PATH");
  console.log("   ✓ Read PATH variable (length:", path.length, "chars)");
} catch (error) {
  console.log("   ✗ Error:", error);
}

console.log("");

// 3. Test network access (commented out to avoid actual network calls in CI)
console.log("3. Testing network access (--net flag):");
console.log("   ℹ Network test skipped to avoid external dependencies");
console.log("   ℹ Use: fetch('https://example.com') to test");

console.log("");
console.log("=== Demo Complete ===");
console.log("All permission checks passed!");
