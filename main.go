package main
import (
	"fmt"
	)

func main() {
	fmt.Println("Hello, World!")
}

func convertToMorse(word string) string	{
	if word == "hola"{
		return ".... --- .-.. .- "
	} else{
		return word
	}
}


