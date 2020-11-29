package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestHandler(t *testing.T) {
	testSuite := []struct {
		name        string
		firstvalue  string
		secondvalue string
		action      string
		answer      float64
		status      int
		err         string
	}{
		{name: "addition", firstvalue: "2", secondvalue: "4", action: "add", answer: 6},
		{name: "subtraction", firstvalue: "2", secondvalue: "4", action: "subtract", answer: -2},
		{name: "multiplication", firstvalue: "2", secondvalue: "4", action: "multiply", answer: 8},
		{name: "division", firstvalue: "4", secondvalue: "4", action: "divide", answer: 1},
		{name: "missing a value", firstvalue: "", secondvalue: "4", action: "add", err: "missing value"},
		{name: "missing a value", firstvalue: "2", secondvalue: "", action: "add", err: "missing value"},
	}

	for _, testCase := range testSuite {
		t.Run(testCase.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:5050/?first_number="+testCase.firstvalue+"&sec_number="+testCase.secondvalue+"&action="+testCase.action, nil)
			if err != nil {
				t.Fatalf("Cannot create request: %v", err)
			}
			rec := httptest.NewRecorder()
			index(rec, req)

			result := rec.Result()
			defer result.Body.Close()

			body, err := ioutil.ReadAll(result.Body)
			if err != nil {
				t.Fatalf("response body cannot read: %v", err)
			}

			if testCase.err != "" {
				if result.StatusCode != http.StatusBadRequest {
					t.Errorf("expected bad request; got %v", result.StatusCode)
				}
				if msg := string(bytes.TrimSpace(body)); msg != testCase.err {
					t.Errorf("expected message %q but got %q", testCase.err, msg)
				}
				return
			}

			if result.StatusCode != http.StatusOK {
				t.Errorf("expected status OK but got %v", result.Status)
			}

			data, err := strconv.ParseFloat(string(bytes.TrimSpace(body)), 64)
			if err != nil {
				t.Fatalf("expected a float but got %s", body)
			}

			if data != testCase.answer {
				t.Fatalf("expected result of %v but got %v", testCase.answer, data)
			}
		})
	}
}

func TestServerRouting(t *testing.T) {
	testSuite := []struct {
		name        string
		firstvalue  string
		secondvalue string
		action      string
		answer      float64
		status      int
		err         string
	}{
		{name: "addition", firstvalue: "2", secondvalue: "4", action: "add", answer: 6},
		{name: "subtraction", firstvalue: "2", secondvalue: "4", action: "subtract", answer: -2},
		{name: "multiplication", firstvalue: "2", secondvalue: "4", action: "multiply", answer: 8},
		{name: "division", firstvalue: "4", secondvalue: "4", action: "divide", answer: 1},
		{name: "missing a value", firstvalue: "", secondvalue: "4", action: "add", err: "missing value"},
		{name: "missing a value", firstvalue: "2", secondvalue: "", action: "add", err: "missing value"},
	}

	server := httptest.NewServer(handlers())
	defer server.Close()
	for _, testCase := range testSuite {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := http.Get(fmt.Sprintf("%s/?first_number="+testCase.firstvalue+"&sec_number="+testCase.secondvalue+"&action="+testCase.action, server.URL))
			if err != nil {
				t.Fatalf("Get request error: %v", err)
			}
			defer result.Body.Close()

			body, err := ioutil.ReadAll(result.Body)
			if err != nil {
				t.Fatalf("response body cannot read: %v", err)
			}

			if testCase.err != "" {
				if result.StatusCode != http.StatusBadRequest {
					t.Errorf("expected bad request; got %v", result.StatusCode)
				}
				if msg := string(bytes.TrimSpace(body)); msg != testCase.err {
					t.Errorf("expected message %q but got %q", testCase.err, msg)
				}
				return
			}

			if result.StatusCode != http.StatusOK {
				t.Errorf("expected status OK but got %v", result.Status)
			}

			data, err := strconv.ParseFloat(string(bytes.TrimSpace(body)), 64)
			if err != nil {
				t.Fatalf("expected a float but got %s", body)
			}

			if data != testCase.answer {
				t.Fatalf("expected result of %v but got %v", testCase.answer, data)
			}
		})
	}
}
