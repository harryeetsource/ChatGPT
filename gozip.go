package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mkobetich/go-zip"
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		inputFilePath := ui.NewLineEdit()
		outputFilePath := ui.NewLineEdit()

		zipButton := ui.NewButton("Zip")
		unzipButton := ui.NewButton("Unzip")

		form := ui.NewForm()
		form.Add("Input File Path", inputFilePath)
		form.Add("Output File Path", outputFilePath)
		form.Add("", zipButton)
		form.Add("", unzipButton)

		window := ui.NewWindow("Zip/Unzip Utility", 400, 120, form)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		zipButton.OnClicked(func(*ui.Button) {
			if inputFilePath.Text() == "" || outputFilePath.Text() == "" {
				log.Fatalln("Input/output file path can't be empty")
			}
			zipFile(inputFilePath.Text(), outputFilePath.Text())
		})

		unzipButton.OnClicked(func(*ui.Button) {
			if inputFilePath.Text() == "" || outputFilePath.Text() == "" {
				log.Fatalln("Input/output file path can't be empty")
			}
			unzipFile(inputFilePath.Text(), outputFilePath.Text())
		})

		window.Show()
	})

	if err != nil {
		log.Fatalf("failed to start GUI: %v", err)
	}
}

func zipFile(inputFilePath string, outputFilePath string) {
	fmt.Println("Zipping file...")

	zipWriter, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Failed to create zip file: %v", err)
	}
	defer zipWriter.Close()

	zipArchiver := zip.NewWriter(zipWriter)
	defer zipArchiver.Close()

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer inputFile.Close()

	info, err := inputFile.Stat()
	if err != nil {
		log.Fatalf("Failed to get file information: %v", err)
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		log.Fatalf("Failed to get file header: %v", err)
	}

	header.Name
