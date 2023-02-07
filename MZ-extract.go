package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

const (
	PEHeader = 0x50450000
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run extract_pe.go <dump_file>")
		os.Exit(1)
	}

	dumpFile := os.Args[1]

	f, err := os.Open(dumpFile)
	if err != nil {
		fmt.Println("Error opening dump file:", err)
		os.Exit(1)
	}
	defer f.Close()

	var peHeader uint32
	err = binary.Read(f, binary.LittleEndian, &peHeader)
	if err != nil {
		fmt.Println("Error reading PE header:", err)
		os.Exit(1)
	}

	if peHeader != PEHeader {
		fmt.Println("Error: Not a valid PE file.")
		os.Exit(1)
	}

	var peSize uint32
	err = binary.Read(f, binary.LittleEndian, &peSize)
	if err != nil {
		fmt.Println("Error reading PE size:", err)
		os.Exit(1)
	}

	peBytes := make([]byte, peSize)
	_, err = f.Read(peBytes)
	if err != nil {
		fmt.Println("Error reading PE bytes:", err)
		os.Exit(1)
	}

	peFile := "pe_extracted.exe"
	pef, err := os.Create(peFile)
	if err != nil {
		fmt.Println("Error creating PE file:", err)
		os.Exit(1)
	}
	defer pef.Close()

	_, err = pef.Write(peBytes)
	if err != nil {
		fmt.Println("Error writing PE bytes:", err)
		os.Exit(1)
	}

	fmt.Printf("PE file extracted successfully to %s\n", peFile)
}
