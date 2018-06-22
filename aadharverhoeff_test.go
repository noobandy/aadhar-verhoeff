package aadharverhoeff

import (
	"testing"
)


func TestValidateVerhoeff(t *testing.T) {
	testCases := [] struct {
		aadhar string
		valid bool
	} {
		{"788198737538", true},
		{"012345678912", false},
		{"01234X678912", false},
	}

	for _, testCase := range testCases {
		if result := ValidateVerhoeff(testCase.aadhar); result != testCase.valid {
			t.Errorf("failed for %v, expected %v got %v\n", testCase.aadhar, testCase.valid, result)
		}
	}

}


func BenchmarkValidateVerhoeff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValidateVerhoeff("788198737538")
	}

}



