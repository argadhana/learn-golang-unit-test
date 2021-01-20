package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//Table Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Arga",
			request: "Arga",
		},
		{
			name:    "Affan",
			request: "Affan",
		},
		{
			name:    "Argadhana",
			request: "Argadhana",
		},
		{
			name:    "Affan Ghifari",
			request: "Affan Ghifari",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

//Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Arga", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Arga")
		}
	})
	b.Run("Affan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Affan")
		}
	})
}

//Table Test
func TestTableHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Arga",
			request:  "Arga",
			expected: "Hello Arga",
		},

		{
			name:     "Affan",
			request:  "Affan",
			expected: "Hello Affan",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

//SUB TESTING
func TestSubTest(t *testing.T) {
	t.Run("Arga", func(t *testing.T) {
		result := HelloWorld("Arga")
		assert.Equal(t, "Hello Arga", result, "Harusnya Hello Arga")
	})

	t.Run("Affan", func(t *testing.T) {
		result := HelloWorld("Affan")
		assert.Equal(t, "Hello Affan", result, "Harusnya Hello Affan")
	})
}

//Main Test (After Before Test)
func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")
	m.Run()
	fmt.Println("Sesuda Unit Test")
}

//Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("cannot run in windows")
	}

	result := HelloWorld("Arga")
	require.Equal(t, "Hello Arga", result, "Harusnya Hello Arga")
}

//Assert Test
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Arga")
	assert.Equal(t, "Hello Arga", result, "Harusnya Hello Arga")
	fmt.Println("Test with Assertion Done")
}

//Require Test
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Arga")
	require.Equal(t, "Hello Arga", result, "Harusnya Hello Arga")
	fmt.Println("Test with Require Done")
}

//Testing
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Arga")

	if result != "Hello Arga" {
		t.Error("Harusnya Hello Arga")
	}

	fmt.Println("Test Hello World Done")
}

func TestHelloWorldAffan(t *testing.T) {
	result := HelloWorld("Affan")

	if result != "Hello Affan" {
		t.Fatal("Harusnya Hello Affan")
	}

	fmt.Println("Test Hello World Affan Done")

}
