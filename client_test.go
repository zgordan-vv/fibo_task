package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const baseURL = "http://127.0.0.1:8080"

// This test checks if a Fibonacci sequence produced by the API is correct for the first 50 elements
func TestFunctionality(t *testing.T) {
	makeRequest(baseURL + "/reset")
	results := []uint64{}
	getURL := baseURL + "/next"
	for i := 0; i < 52; i++ {
		if i == 14 || i == 37 {
			makeRequest(baseURL + "/make_trouble")
			continue
		}
		b, err := makeRequest(getURL)
		if err != nil {
			t.Fatal(err)
		}
		value, err := strconv.ParseUint(string(b), 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		results = append(results, value)
	}
	for i := 0; i < 50; i++ {
		fmt.Println(i, results[i])
	}
	assert.Equal(t, results[49], uint64(12586269025))
}

// This test checks if 1000 requests take more than 1000 ms
func TestPerformance(t *testing.T) {
	getURL := baseURL + "/current"
	start := time.Now().UnixMilli()
	for i := 0; i < 1000; i++ {
		if _, err := makeRequest(getURL); err != nil {
			t.Fatal(err)
		}
	}
	end := time.Now().UnixMilli()
	timeToRun := end - start
	if end-start > 1000 { // see if the total duration exceeds 1000 ms
		t.Errorf("Requests took too long: %d", timeToRun)
	}
	fmt.Println("Total execution time:", timeToRun)
}

func makeRequest(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf(fmt.Sprintf("Unexpected status code: %d", resp.StatusCode))
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
