package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Use the `ps` command to get a list of processes
	out, err := exec.Command("ps", "-e", "-opid,comm").Output()
	if err != nil {
		log.Fatal(err)
	}

	// Print the output of the `ps` command
	fmt.Printf("%s", out)
}
