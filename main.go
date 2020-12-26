package main
import (
	"fmt"
	"strings"
	)

var morseDicctionaty = map[string]string{
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
	"?": "..--..",
	"!": "-.-.--",
	"-": "-....-",
	"/": "-..-.",
	"@": ".--.-.",
	"(": "-.--.",
	")": "-.--.-",
}

func toMorseCode(word, splitter string) string {
	var morse string
	word = strings.ToUpper(word)

	for i:= 0 ; i< len(word); i++{
		code := morseDicctionaty[word[i:i+1]]
		if code != ""{
			morse += code + splitter
		}
	}
	return morse
}

func main() {
	fmt.Println("Hello, World!")
}

func convertToMorse(word string) string	{
	var splitter = " "
	return toMorseCode(word, splitter)
}


