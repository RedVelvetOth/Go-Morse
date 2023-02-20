package morse

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myMap = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",

	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	"0": "-----",

	".": ".-.-.-",
	",": "--..--",
	":": "---...",
	"?": "..--..",
	"-": "-....-",
	"/": "-..-.",
	"(": "-.--.",
	")": "-.--.-",
	"“": ".-..-.",
	"+": ".-.-.",
	"@": ".--.-.",
	" ": "/",
}

var reverseMap = make(map[string]string)

func Morser(s string) (v string) {
	for _, x := range strings.ToUpper(s) {
		v = v + " " + myMap[string(x)]
	}
	return
}

func reversingMap() {
	for key, value := range myMap {
		reverseMap[value] = key
	}
}

func Demorser(s string) (v string) {
	reversingMap()
	for _, x := range strings.Fields(s) {
		v = v + reverseMap[string(x)]
	}
	return
}

func Gui() {
	myapp := app.New()
	mywindow := myapp.NewWindow("Morser Or Demorser")
	icon, _ := fyne.LoadResourceFromURLString("https://cdn-icons-png.flaticon.com/512/260/260352.png")
	mywindow.SetIcon(icon)

	input := widget.NewEntry()
	input.SetPlaceHolder("Do What Ya Wanna do")

	button_m := widget.NewButtonWithIcon("Morser", theme.ContentCopyIcon(), func() {
		text := Morser(input.Text)
		input.SetText(text)
		mywindow.Clipboard().SetContent(text)
	})

	button_dm := widget.NewButtonWithIcon("Demorser", theme.ContentCopyIcon(), func() {
		text := Demorser(input.Text)
		input.SetText(text)
		mywindow.Clipboard().SetContent(text)
	})

	mywindow.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				input,
				nil,
				nil,
				nil,
			),
			input,
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(2),
				button_m,
				button_dm,
			),
		),
	)
	mywindow.Resize(fyne.NewSize(400, 50))
	mywindow.ShowAndRun()

}
