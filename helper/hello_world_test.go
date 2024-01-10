package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World")

	if result != "Hello World" {
		// error
		panic("Result is not 'Hello World' ")
	}
}
