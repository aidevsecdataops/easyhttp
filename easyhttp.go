package easyhttp

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Hello prints a string for testing
func Hello() {
	log.Println("Exported Hello")
}

const (
	timeoutInSecs = 60
)

// Get makes an HTTP Get call
func Get(url string) ([]byte, int, error) {
	timeout := time.Duration(timeoutInSecs * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-type", "application/json")

	if err != nil {
		return nil, 0, err
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, 0, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil

}

// Post makes an HTTP Post call
func Post(url string, requestBody []byte) ([]byte, int, error) {
	// We will set our default timeout
	timeout := time.Duration(timeoutInSecs * time.Second)

	// We create a client with the timeout attached
	client := http.Client{
		Timeout: timeout,
	}

	// We are going to create a new request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	// We set the header to application/json
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		return nil, 0, err
	}
	// We execute our request
	resp, err := client.Do(request)
	if err != nil {
		return nil, 0, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	// We parse our response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	// We return a byte array
	return body, resp.StatusCode, nil
}

// Put makes an HTTP Put call
func Put(url string, requestBody []byte) ([]byte, *http.Response, error) {
	// We will set our default timeout
	timeout := time.Duration(timeoutInSecs * time.Second)

	// We create a client with the timeout attached
	client := http.Client{
		Timeout: timeout,
	}

	// We are going to create a new request
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))

	// We set the header to application/json
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		return nil, nil, err
	}
	// We execute our request
	resp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	// We parse our response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	// We return a byte array
	return body, resp, nil
}
