package main

import 
	"testing"


func TestConvertion1(t *testing.T){
	if convertToMorse("hola") != ".... --- .-.. .- "  {
		t.Error("wrong case 1")
	}
}
func TestConvertion2(t *testing.T){
	if convertToMorse("tucan") != "..- -.-. .- -. "  {
		t.Error("wrong case 2")
	}
}