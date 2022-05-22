package utility

import (
	"io/ioutil"
	"net/http"
)

// ReadBody returns body of a request as string
func ReadBody(r *http.Request) (string, error) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(bytes[:]), nil
}

// ReadBody returns body of a request as string
func ReadBodyAsBytes(r *http.Request) ([]byte, error) {
	bytes, err := ioutil.ReadAll(r.Body)
	return bytes, err
}
