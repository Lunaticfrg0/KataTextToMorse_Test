package main

import 
	"testing"


func TestConvertion1(t *testing.T){
	if convertToMorse("hola") != ".... --- .-.. .- "  {
		t.Error("wrong case 1")
	}
}
