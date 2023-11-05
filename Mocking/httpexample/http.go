package httpexample

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrBadStatusCode = errors.New("bad request status code from HTTP call")

type Response struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func MakeHttpCall(url string) (*Response, error) {
	resp, err := http.Get(fmt.Sprintf("%s/resource/55/foo", url))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrBadStatusCode
	}
	defer resp.Body.Close()
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
