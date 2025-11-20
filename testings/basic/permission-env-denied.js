// Test that environment variable operations are denied without --env flag
try {
  const value = Kimera.getEnv("PATH");
  console.log("ERROR: Should have denied environment access");
  close();
} catch (error) {
  if (error.toString().includes("environment variable access denied")) {
    console.log("PASS: Environment access correctly denied");
  } else {
    console.log("ERROR: Wrong error message:", error);
    close();
  }
}
