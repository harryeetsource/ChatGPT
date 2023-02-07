package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	openProcess := kernel32.NewProc("OpenProcess")
	readProcessMemory := kernel32.NewProc("ReadProcessMemory")
	closeHandle := kernel32.NewProc("CloseHandle")

	var (
		processHandle syscall.Handle
		readBytes uint32
	)

	// Replace PID with the desired process ID
	pid := uint32(1234)

	// Open the process
	r1, _, _ := openProcess.Call(uintptr(0x0010), uintptr(0), uintptr(pid))
	processHandle = syscall.Handle(r1)
	if processHandle == 0 {
		fmt.Println("Failed to open process")
		return
	}

	// Allocate memory to store the memory dump
	buffer := make([]byte, 1024)

	// Replace address with the desired memory address to dump
	address := uintptr(0x00000000)

	// Read the memory
	r2, _, _ := readProcessMemory.Call(uintptr(processHandle), uintptr(address), uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)), uintptr(unsafe.Pointer(&readBytes)))
	if r2 == 0 {
		fmt.Println("Failed to read memory")
		return
	}

	// Close the process handle
	r3, _, _ := closeHandle.Call(uintptr(processHandle))
	if r3 == 0 {
		fmt.Println("Failed to close handle")
		return
	}

	fmt.Printf("Memory dump: %x", buffer)
}
