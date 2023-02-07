package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	// Get a handle to the Windows system process and duplicate it for our usage
	h, err := windows.GetCurrentProcess()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer windows.CloseHandle(h)

	var snap windows.Handle
	snap, err = windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer windows.CloseHandle(snap)

	var pe32 windows.ProcessEntry32
