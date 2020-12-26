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
func TestConvertion5(t *testing.T){
	if convertToMorse("camaleon") != "-.-. .- -- .- .-.. . --- -. "  {
		t.Error("wrong case 5")
	}
}
func TestConvertion6(t *testing.T){
	if convertToMorse("(ayuda)") != "-.--. .- -.-- ..- -.. .- -.--.- "  {
		t.Error("wrong case 6")
	}
}