package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Define the source directory to backup
	srcDir := "/path/to/source/directory"

	// Define the destination directory for the backup
	dstDir := "/path/to/destination/directory"

	// Use the `tar` command to create a gzipped archive of the source directory
	tarCmd := exec.Command("tar", "-czf", dstDir+"/backup.tar.gz", "-C", srcDir, ".")

	// Run the `tar` command and capture the output
	output, err := tarCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error creating backup: %s\n", err)
		os.Exit(1)
	}

	// Print the output of the `tar` command
	fmt.Printf("Backup created successfully: %s\n", output)
}
