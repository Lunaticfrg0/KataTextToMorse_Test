package main

import 
	"testing"


func TestConvertion1(t *testing.T){
	if convertToMorse("hola") != ".... --- .-.. .- "  {
		t.Error("wrong case 1")
	}
}
func TestConvertion2(t *testing.T){
	if convertToMorse("tucan") != "- ..- -.-. .- -. "  {
		t.Error("wrong case 2")
	}
}
func TestConvertion3(t *testing.T){
	if convertToMorse("1233") != ".---- ..--- ...-- ...-- "  {
		t.Error("wrong case 3")
	}
}
func TestConvertion4(t *testing.T){
	if convertToMorse("SOS") != "... --- ... "  {
		t.Error("wrong case 4")
	}
}