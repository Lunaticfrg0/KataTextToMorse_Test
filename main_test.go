package main

import 
	"testing"


func TestConvertion1(t *testing.T){
	if convertion("hola") != ".... --- .-.. .- "{
		t.Error("wrong")
	}
}
func TestConvertion2(t *testing.T){
	if convertion("queso") != "--.- ..- . ... --- "{
		t.Error("wrong")
	}
}

