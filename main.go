package main
import (
	"fmt"
	)

func main() {
	fmt.Println("Hello, World!")
}

func convertionToMorse(word string) string{
	if word == "hola"{
		return ".... --- .-.. .- "
	}else {
		return word
	}
	
}
