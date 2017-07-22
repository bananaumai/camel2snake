package converter

import "testing"

func TestConvert(t *testing.T) {
	testCases := []struct {
		test string
		expect string
	} {
		{"CamelCase", "camel_case"},
		{"lowerCamelCase", "lower_camel_case"},
		{"EndWith1", "end_with1"},
	}

	for _, testCase := range testCases {
		if Convert(testCase.test) != testCase.expect {
			t.Fatalf("Failed in test case : %v\n", testCase)
		}
	}
}
