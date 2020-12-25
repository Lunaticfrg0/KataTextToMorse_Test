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

// func toMorseCode(text, spliter string)string{
// 	var output string
// 	text = strings.
// }

func toMorseCodeWord(text, spliter string) string{
	var morse string
	text = strings.ToUpper(text)
	for i:=0 ; i<len(text); i++{
		code := morseDicctionary[text[i:i+1]]
		if code != ""{
			morse+=code+spliter
		}
	}
	return morse
}


func main() {
	fmt.Println("Hello, World!")
}

func convertion(message string) string {
		
	return toMorseCodeWord(message, " ")
}
