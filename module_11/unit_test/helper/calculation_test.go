package unit_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ! Untuk menjalankan testing go test -v, atau go test -v -cover untuk melihat coverage

func TestFailSum(t *testing.T) {
	result := Sum(1, 2)

	if result != 4 {
		// t.Fail() menggagalkan test
		// t.FailNow() apabila menggunakan ini, fmt.Println tidak akan dijalankan
		// t.Error("Result has to be 4") ini akan menggagalkan test dan logging error
		t.Fatal("Result has to be 4") // ini akan menggagalkan test seperti FailNow() dan logging error
	}
	fmt.Println("TestFailSum() eksekusi terhenti")
}

func TestSum(t *testing.T) {
	result := Sum(1, 2)

	if result != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", result, 3)
	}
}

// * Cara Proper menggunakan testify untuk assert dan require, bedanya assert tidak akan menghentikan eksekusi test, sedangkan require akan menghentikan eksekusi test saat terjadi error
func TestFailSum2(t *testing.T) {
	expectedResult := 20
	result := Sum(5, 5)

	// Require (interface testingT, expected, actual, message)
	require.Equal(t, expectedResult, result, "Result has to be 20")
	fmt.Println("TestFailSum2() eksekusi terhenti") // ini tidak tereksekusi karena require
}

func TestSum2(t *testing.T) {
	expectedResult := 20
	result := Sum(10, 10)

	//  Assert (interface testingT, expected, actual, message)
	assert.Equal(t, expectedResult, result, "Result has to be 20")

	fmt.Println("TestSum2() eksekusi terhenti")
}

func TestTableSum(t *testing.T) {
	tests := []struct {
		request    int
		expected   int
		errMessage string
	}{
		{
			request:    Sum(1, 2),
			expected:   3,
			errMessage: "Result has to be 3",
		},
		{
			request:    Sum(5, 5),
			expected:   10,
			errMessage: "Result has to be 10",
		},
		{
			request:    Sum(10, 10),
			expected:   20,
			errMessage: "Result has to be 20",
		},
	}

	for i, test := range tests {
		// Menjalankan sub test (string, callback func)
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, test.expected, test.request, test.errMessage)
		})
	}

}
