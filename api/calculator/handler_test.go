package calculator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestSumRequest(t *testing.T) {
	testData := []struct {
		name           string
		requestData    string
		expectedResult int
		status         int
		err            string
	}{
		{name: "numbers from 1 to 4", requestData: "numbers=1&numbers=2&numbers=3&numbers=4", expectedResult: 10, status: http.StatusOK},
		{name: "no numbers sent", requestData: "", expectedResult: 0, status: http.StatusOK},
		{name: "character instead of number sent", requestData: "numbers=a", err: "No numbers or special characters allowed."},
	}

	for _, testCase := range testData {
		t.Run(testCase.name, func(t *testing.T) {
			request, err := http.NewRequest("GET",
				fmt.Sprintf("localhost:8888/calculator/sum-integers?%s", testCase.requestData),
				nil)
			if err != nil {
				t.Fatalf("could not calculate sum of %s: %v", testCase.name, err)
			}

			recorder := httptest.NewRecorder()

			SumIntegerHandler(recorder, request)

			response := recorder.Result()
			defer response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Fatalf("could not read response body for sum of %s: %v", testCase.name, err)
			}

			if testCase.err != "" {
				if response.StatusCode != http.StatusBadRequest {
					t.Errorf("expected Bad requets status, but got %v", response.StatusCode)
				}

				returnedErrormessage := string(bytes.TrimSpace(body))
				if returnedErrormessage != testCase.err {
					t.Errorf("Wrong error message returned, expected %s, got %s", testCase.err, returnedErrormessage)
				}
				return
			}

			if response.StatusCode != http.StatusOK {
				t.Errorf("expected status OK, but got %v", response.StatusCode)
			}

			result, err := strconv.Atoi(string(body))
			if err != nil {
				t.Fatalf("could not read convert response body to correct format for sum of %s: %v", testCase.name, err)
			}

			if result != testCase.expectedResult {
				t.Fatalf("Incorect result for sum of %s, got %d, expected %d", testCase.name, result, testCase.expectedResult)
			}
		})
	}
}
