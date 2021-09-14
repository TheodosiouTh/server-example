package calculator

import "testing"

func TestSum(t *testing.T) {
	testData := []struct {
		name           string
		numbers        []int
		expectedResult int
	}{
		{"nil", nil, 0},
		{"empty array", []int{}, 0},
		{"numbers from 1 to 4", []int{1, 2, 3, 4}, 10},
	}

	for _, testCase := range testData {
		receivedResult := Sum(testCase.numbers)
		if receivedResult != testCase.expectedResult {
			t.Errorf("Sum of %s failed, got %d, expected %d", testCase.name, receivedResult, testCase.expectedResult)
		}
	}
}
