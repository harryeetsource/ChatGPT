package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	createToolhelp32Snapshot := kernel32.NewProc("CreateToolhelp32Snapshot")
	process32First := kernel32.NewProc("Process32FirstW")
	process32Next := kernel32.NewProc("Process32NextW")
	openProcess := kernel32.NewProc("OpenProcess")
	readProcessMemory := kernel32.NewProc("ReadProcessMemory")

	var snapshot syscall.Handle
	ret, _, _ := createToolhelp32Snapshot.Call(uintptr(
