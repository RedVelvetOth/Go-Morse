package main

import (
	"os"
	"toys/morse/morse"
)

func main() {
	os.Setenv("FYNE_THEME", "dark")

	morse.Gui()
}
