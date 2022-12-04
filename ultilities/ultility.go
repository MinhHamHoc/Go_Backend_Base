package utilities

import (
	"errors"
	"net/http"

	"github.com/joho/godotenv"
)

func LoadEnvFromFile(envPath string) (err error) {
	err = godotenv.Load(envPath)
	if err != nil {
		return errors.New("cannot get env")
	}
	return nil
}

func GetQuery(req *http.Request, key string) (string, bool) {
	if values := req.URL.Query().Get(key); len(values) > 0 {
		return values, true
	}
	return "", false
}

func IntInArray(key int, arr []int) bool {
	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2

		if arr[median] < key {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != key {
		return false
	}

	return true
}
