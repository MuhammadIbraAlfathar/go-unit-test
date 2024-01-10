package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldSAtu(t *testing.T) {
	result := HelloWorld("World")

	if result != "Hello World" {
		// error
		t.Error("Result must be 'Hello World'")
	}

	fmt.Println("TestHelloWorldSatu done")

}

func TestHelloWorldDua(t *testing.T) {
	result := HelloWorld("World")

	if result != "Hello World" {
		// error
		t.Fatal("Result must be 'Hello World'")
	}

	fmt.Println("TestHelloWorldDua done")
}
