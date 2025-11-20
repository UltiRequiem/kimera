// Test that filesystem operations work with --fs flag
try {
  const content = Kimera.readFile("readme.md");
  if (content.includes("Kimera")) {
    console.log("PASS: Filesystem access works with --fs flag");
  } else {
    console.log("ERROR: Could not read file correctly");
    close();
  }
} catch (error) {
  console.log("ERROR: Should have allowed filesystem access:", error);
  close();
}
