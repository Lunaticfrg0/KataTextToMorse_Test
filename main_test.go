package main

import 
	"testing"


func TestConvertion1(t *testing.T){
	if convertion("hola") != ".... --- .-.. .- "{
		t.Error("wrong")
	}
}
