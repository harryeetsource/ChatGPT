package main

import (
	"bufio"
	"fmt"
	"github.com/ulikunitz/xz"
	"os"
)

func main() {
	fmt.Println("1. Compress")
	fmt.Println("2. Decompress")
	fmt.Print("Choose an option: ")

	var option int
	fmt.Scanf("%d\n", &option)

	switch option {
	case 1:
		compress()
	case 2:
		decompress()
	default:
		fmt.Println("Invalid option")
	}
}

func compress() {
	var fileIn, fileOut string
	fmt.Print("Enter input file name: ")
	fmt.Scanf("%s\n", &fileIn)
	fmt.Print("Enter output file name: ")
	fmt.Scanf("%s\n", &fileOut)

	fin, err := os.Open(fileIn)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer fin.Close()

	fout, err := os.Create(fileOut)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer fout.Close()

	w := xz.NewWriter(fout)
	defer w.Close()

	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		_, err := w.Write(scanner.Bytes())
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	fmt.Println("Compression completed")
}

func decompress() {
	var fileIn, fileOut string
	fmt.Print("Enter input file name: ")
	fmt.Scanf("%s\n", &fileIn)
	fmt.Print("Enter output file name: ")
	fmt.Scanf("%s\n", &fileOut)

	fin, err := os.Open(fileIn)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer fin.Close()

	fout, err := os.Create(fileOut)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer fout.Close()

	r, err := xz.NewReader(fin, 0)
	if err != nil {
		fmt.Println("Error opening xz reader:", err)
		return
	}

	
