package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Benchmark

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("ibra")
	}
}

func BenchmarkHelloWorld2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Worlddd")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Test 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Worlddd")
		}
	})

	b.Run("Test 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ibra")
		}
	})

}

func BenchmarkTableTest(b *testing.B) {
	benchmarks := []struct {
		Name    string
		Request string
	}{
		{
			Name:    "Test 1",
			Request: "World",
		},
		{
			Name:    "Test 2",
			Request: "Eko",
		},
		{
			Name:    "Test 3",
			Request: "Kurniawan",
		},
		{
			Name:    "Test 4",
			Request: "Jeki",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.Request)
			}
		})
	}
}

// func TestMain(m *testing.M) {
// 	// before
// 	fmt.Println("BEFORE UNIT TEST")

// 	m.Run()

// 	//after
// 	fmt.Println("AFTER UNIT TEST")
// }

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

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		Name     string
		Request  string
		Expected string
	}{
		{
			Name:     "Test 1",
			Request:  "World",
			Expected: "Hello World",
		},
		{
			Name:     "Test 2",
			Request:  "Eko",
			Expected: "Hello Eko",
		},
		{
			Name:     "Test 3",
			Request:  "Joko",
			Expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := HelloWorld(test.Request)
			assert.Equal(t, test.Expected, result)
		})
	}

}
