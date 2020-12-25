package main
import (
	"fmt"
	"strings"
	)
var morseDicctionary = map[string]string{
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


func main() {
	fmt.Println("Hello, World!")
}

func toMorseCode(word, splitter string) string {
	var morse string
	word = strings.ToUpper(word)
	for i:= 0 ; i < len(word); i++{
		code:= morseDicctionary[word[i:i+1]]
		if code != ""{
			morse += code+splitter
		}
	}
	return morse
}

func convertionToMorse(word string) string{
	splitter := " "
	return toMorseCode(word, splitter)
	
}
