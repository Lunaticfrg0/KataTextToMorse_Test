package main

import 
	"testing"


func TestConvertion1(t *testing.T){
	if convertionToMorse("hola") != ".... --- .-.. .- " {
		t.Error("wrong")
	}
}
func TestConvertion2(t *testing.T){
	if convertionToMorse("test") != "- . ... - " {
		t.Error("wrong")
	}
}

