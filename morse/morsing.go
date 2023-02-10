package morse

import (
	"strings"
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
	for _, x := range s {
		v = v + reverseMap[string(x)]
	}
	return
}
