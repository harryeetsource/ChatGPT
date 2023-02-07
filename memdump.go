package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid := os.Getpid()
	proc, _ := os.FindProcess(pid)
	f, _ := os.Create("memory.dump")
	defer f.Close()
	p, _ := proc.Memory()
	f.Write(p)
	fmt.Println("Dumped process memory to disk.")
}
