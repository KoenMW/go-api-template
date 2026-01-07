package rest

import (
	"encoding/json"
	"errors"
	"go-api/domain/core"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func BodyReader[T any](req *http.Request) (T, error) {
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		var zero T
		return zero, errors.New(core.UnableToReadBody)
	}

	var data T
	if err := json.Unmarshal(body, &data); err != nil {
		var zero T
		return zero, err
	}

	return data, nil
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

func ParsePathID(r *http.Request) (uuid.UUID, error) {
	idStr := r.PathValue("id")
	if idStr == "" {
		return uuid.Nil, errors.New(core.InvalidId)
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, errors.New(core.InvalidId)
	}

	return id, nil
}

func CORSHandler(next http.Handler) http.Handler {
	env := os.Getenv("ENV")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if env == "dev" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		}

		if strings.HasSuffix(origin, ".github.io") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
