// Test that multiple permission flags work together
try {
  // Test filesystem
  const content = Kimera.readFile("readme.md");
  if (!content.includes("Kimera")) {
    console.log("ERROR: Filesystem access failed");
    close();
  }
  
  // Test environment
  const path = Kimera.getEnv("PATH");
  if (!path || path.length === 0) {
    console.log("ERROR: Environment access failed");
    close();
  }
  
  console.log("PASS: All permissions work together with multiple flags");
} catch (error) {
  console.log("ERROR: Combined permissions test failed:", error);
  close();
}
