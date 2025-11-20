// Test that file write operations are denied without --fs flag
try {
  Kimera.writeFile("/tmp/kimera-test.txt", "test content");
  console.log("ERROR: Should have denied filesystem access");
  close();
} catch (error) {
  if (error.toString().includes("filesystem access denied")) {
    console.log("PASS: File write correctly denied");
  } else {
    console.log("ERROR: Wrong error message:", error);
    close();
  }
}
