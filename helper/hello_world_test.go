package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		result := HelloWorld("World")

		assert.Equal(t, "Hello World", result, "Result mus be 'Hello World'")
	})

	t.Run("Test 2", func(t *testing.T) {
		result := HelloWorld("World")

		assert.Equal(t, "Hello World", result, "Result mus be 'Hello World'")
	})
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("World")

	assert.Equal(t, "Hello World", result, "Result mus be 'Hello World'")

	fmt.Println("TestHelloWorldAssert is done")
}

func TestHelloWorldSatu(t *testing.T) {
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
