package main

import 
	"testing"


func TestConvertion1(t *testing.T){
	if convertionToMorse("hola") != ".... --- .-.. .- " {
		t.Error("wrong case 1")
	}
}
func TestConvertion2(t *testing.T){
	if convertionToMorse("test") != "- . ... - " {
		t.Error("wrong case 2")
	}
}
func TestConvertion3(t *testing.T){
	if convertionToMorse("cocina") != "-.-. --- -.-. .. -. .- " {
		t.Error("wrong case 3")
	}
}
func TestConvertion4(t *testing.T){
	if convertionToMorse("queso") != "--.- ..- . ... --- " {
		t.Error("wrong case 4")
	}
}
func TestConvertion5(t *testing.T){
	if convertionToMorse("camaleon") != "-.-. .- -- .- .-.. . --- -. " {
		t.Error("wrong case 5")
	}
}
func TestConvertion6(t *testing.T){
	if convertionToMorse("hp") != ".... .--. " {
		t.Error("wrong case 6")
	}
}
func TestConvertion7(t *testing.T){
	if convertionToMorse("1234567") != ".---- ..--- ...-- ....- ..... -.... --... " {
		t.Error("wrong case 6")
	}
}