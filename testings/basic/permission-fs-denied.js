// Test that filesystem operations are denied without --fs flag
try {
  const content = Kimera.readFile("readme.md");
  console.log("ERROR: Should have denied filesystem access");
  close();
} catch (error) {
  if (error.toString().includes("filesystem access denied")) {
    console.log("PASS: Filesystem access correctly denied");
  } else {
    console.log("ERROR: Wrong error message:", error);
    close();
  }
}
