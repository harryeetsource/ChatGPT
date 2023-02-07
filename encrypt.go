package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/andlabs/ui"
)

func hashData(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func main() {
	err := ui.Main(func() {
		input := ui.NewEntry()
		button := ui.NewButton("Encrypt")
		output := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter data to encrypt"), false)
		box.Append(input, false)
		box.Append(button, false)
		box.Append(output, false)
		window := ui.NewWindow("SHA-256 Encryption", 200, 100, false)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			output.SetText(hashData(input.Text()))
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
