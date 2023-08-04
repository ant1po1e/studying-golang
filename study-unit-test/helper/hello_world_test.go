package helper

import (
	"testing"
	"fmt"
	"runtime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)



////////////////////////////// Benchmark
func Benchmark(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		HelloWorld("Antipole")
	}
}



////////////////////////////// Sub Benchmark
func BenchmarkSub(bench *testing.B) {
	bench.Run("Antipole", func (bench *testing.B) {
		for i := 0; i < bench.N; i++ {
			HelloWorld("Antipole")
		}
	})

	bench.Run("Rushia", func (bench *testing.B) {
		for i := 0; i < bench.N; i++ {
			HelloWorld("Rushia")
		}
	})
}



////////////////////////////// Benchmark Table
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name string
		request string
	} {
		{
			name: "Antipole",
			request: "Antipole", 
		},
		{
			name: "Rushia",
			request: "Rushia",
		},
		{
			name: "Azusa",
			request: "Azusa",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func (b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
} 



////////////////////////////// TestMain
func TestMain(main *testing.M) {
	// Before Unit Test
	fmt.Println("---------- Before Unit Test ----------")

	main.Run()

	fmt.Println("---------- After Unit Test ----------")
	// After Unit Test
}



////////////////////////////// Panic()
func TestPanic(test *testing.T) {
	result := HelloWorld("Rushia")
	if (result != "Hello Rushia") {
		panic("Result is not Hello Rushia")
	}

	fmt.Println("++++++++++ Testing Done ++++++++++")
}



////////////////////////////// Fail()
func TestFail(test *testing.T) {
	result := HelloWorld("Antipole")
	if (result != "Hello Antipole") {
		test.Fail()
	}
	 
	fmt.Println("++++++++++ Testing Done ++++++++++")
}



////////////////////////////// FailNow()
func TestFailNow(test *testing.T) {
	result := HelloWorld("Azusa")
	if (result != "Hello Azusa") {
		test.FailNow()
	}
	 
	fmt.Println("++++++++++ Testing Done ++++++++++")
}



////////////////////////////// Error()
func TestError(test *testing.T) {
	result := HelloWorld("Amamiya")
	if (result != "Hello Amamiya") {
		test.Error("Result should be Hello Amamiya")
	}
	 
	fmt.Println("++++++++++ Testing Done ++++++++++")
}



////////////////////////////// Fatal()
func TestFatal(test *testing.T) {
	result := HelloWorld("Shion")
	if (result != "Hello Shion") {
		test.Fatal("Result should be Hello Shion")
	}
	 
	fmt.Println("++++++++++ Testing Done ++++++++++")
}



////////////////////////////// Testify Assertion 
func TestAssert(test *testing.T) {
	result := HelloWorld("Stanford")
	assert.Equal(test, "Hello Stanford", result, "Result should be Hello Stanford")

	fmt.Println("++++++++++ Testing Done ++++++++++")
	
}



////////////////////////////// Testify Require
func TestRequire(test *testing.T) {
	result := HelloWorld("Tystiana")
	require.Equal(test, "Hello Tystiana", result, "Result should be Hello Tystiana")

	fmt.Println("++++++++++ Testing Done ++++++++++")
	
}



////////////////////////////// Skip Test
func TestSkip(test *testing.T) {
	if runtime.GOOS == "windows" {
		test.Skip("Cant run on Windows")
	}
	// if runtime.GOOS == "darwin" {
	// 	test.Skip("Cant run on MAC")
	// }

	result := HelloWorld("Kotori")
	require.Equal(test, "Hello Kotori", result, "Result should be Hello Kotori")
}



////////////////////////////// Sub Test
func TestSub(test *testing.T) {
	test.Run("Panic", func(test *testing.T) {
		result := HelloWorld("Rushia")
		if (result != "Hello Rushia") {
			panic("Result is not Hello Rushia")
		}
	})

	test.Run("Fail", func(test *testing.T) {
		result := HelloWorld("Antipole")
		if (result != "Hello Antipole") {
			test.Fail()
		}
	})

	test.Run("FailNow", func(test *testing.T) {
		result := HelloWorld("Azusa")
		if (result != "Hello Azusa") {
			test.FailNow()
		}
	})

	test.Run("Error", func(test *testing.T) {
		result := HelloWorld("Amamiya")
		if (result != "Hello Amamiya") {
			test.Error("Result should be Hello Amamiya")
		}
	})
	
	test.Run("Fatal", func(test *testing.T) {
		result := HelloWorld("Shion")
		if (result != "Hello Shion") {
			test.Fatal("Result should be Hello Shion")
		}
	})

	test.Run("Assert", func(test *testing.T) {
		result := HelloWorld("Stanford")
		assert.Equal(test, "Hello Stanford", result, "Result should be Hello Stanford")
	})

	test.Run("Require", func(test *testing.T) {
		result := HelloWorld("Stanford")
		require.Equal(test, "Hello Stanford", result, "Result should be Hello Stanford")
	})
}



////////////////////////////// Table Test
func TestTable(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	} {
		{
			name: "Antipole",
			request: "Antipole",
			expected: "Hello Antipole", 
		},
		{
			name: "Rushia",
			request: "Rushia",
			expected: "Hello Rushia", 
		},
		{
			name: "Azusa",
			request: "Azusa",
			expected: "Hello Azusa", 
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
} 
