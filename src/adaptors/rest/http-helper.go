package rest

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func BodyReader[T any](req *http.Request) (*T, error) {
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errors.New("unable to read body")
	}

	var data T
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
