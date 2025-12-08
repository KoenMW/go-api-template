package rest

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
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

func parsePagination(r *http.Request) (int, int) {
	q := r.URL.Query()

	page := 1
	perPage := 10

	if v := q.Get("page"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			page = n
		}
	}

	if v := q.Get("per_page"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			perPage = n
		}
	}

	return page, perPage
}

func WriteJSONResponse[T any](w http.ResponseWriter, statusCode int, data *T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}
